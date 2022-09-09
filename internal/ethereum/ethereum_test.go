package ethereum_test

import (
	"crypto/ecdsa"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"kkn.fi/cmd/ewallet/internal/ethereum"
)

func privateToPublic(t *testing.T, key string) *ecdsa.PublicKey {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		t.Errorf("error while converting hex private key to ECDSA: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Errorf("error casting public key to crypto/ecdsa.PublicKey")
	}
	return publicKeyECDSA
}

func TestPrivateKeyToPublicKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PublicKey
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				key: "f9b05ef4e0fb4e6c11da9df91e45fffa6ba56bde8a192b886ed4d90a57e9ee50",
			},
			want: privateToPublic(t, "f9b05ef4e0fb4e6c11da9df91e45fffa6ba56bde8a192b886ed4d90a57e9ee50"),
		},
		{
			name: "happy path with 0x prefix",
			args: args{
				key: "0xf9b05ef4e0fb4e6c11da9df91e45fffa6ba56bde8a192b886ed4d90a57e9ee50",
			},
			want: privateToPublic(t, "f9b05ef4e0fb4e6c11da9df91e45fffa6ba56bde8a192b886ed4d90a57e9ee50"),
		},
		{
			name: "private key is malformed",
			args: args{
				key: "f91e45fffa6ba56bde8a192b886ed4d90a57e9ee50",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ethereum.PrivateKeyToPublicKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateKeyToPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrivateKeyToPublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivateKeyToAddress(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				key: "f9b05ef4e0fb4e6c11da9df91e45fffa6ba56bde8a192b886ed4d90a57e9ee50",
			},
			want: "0xb7A9E833865e9f2F4dFFED6338DA68130Ed2B319",
		},
		{
			name: "key is malformed",
			args: args{
				key: "f91e45fffa6ba56bde8a192b886ed4d90a57e9ee50",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ethereum.PrivateKeyToAddress(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateKeyToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrivateKeyToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrivateKeyToPublic(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				key: "f9b05ef4e0fb4e6c11da9df91e45fffa6ba56bde8a192b886ed4d90a57e9ee50",
			},
			want: "0488865e02380c8e99d7eb6d79de40757980c81ab35451de6dc7b74cdace024b4b0bdae3e365d489d86c3ba5b739eefff94433ae2fa036fa7ebf114f256bd6f6",
		},
		{
			name: "private key is malformed",
			args: args{
				key: "f91e45fffa6ba56bde8a192b886ed4d90a57e9ee50",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ethereum.PrivateKeyToPublic(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrivateKeyToPublic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrivateKeyToPublic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublicKeyToString(t *testing.T) {
	type args struct {
		publicKeyECDSA *ecdsa.PublicKey
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{
				publicKeyECDSA: privateToPublic(t, "f9b05ef4e0fb4e6c11da9df91e45fffa6ba56bde8a192b886ed4d90a57e9ee50"),
			},
			want: "0488865e02380c8e99d7eb6d79de40757980c81ab35451de6dc7b74cdace024b4b0bdae3e365d489d86c3ba5b739eefff94433ae2fa036fa7ebf114f256bd6f6",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ethereum.PublicKeyToString(tt.args.publicKeyECDSA); got != tt.want {
				t.Errorf("PublicKeyToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressToChecksumCase(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				address: "0xb7a9e833865e9f2f4dffed6338da68130ed2B319",
			},
			want: "0xb7A9E833865e9f2F4dFFED6338DA68130Ed2B319",
		},
		{
			name: "not an address",
			args: args{
				address: "foobar",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ethereum.AddressToChecksumCase(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("ethereum.AddressToChecksumCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ethereum.AddressToChecksumCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
