package server_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/tests"
)

func TestGenerateCoinKey(t *testing.T) {
	t.Parallel()
	addr, mnemonic, err := server.GenerateCoinKey()
	require.NoError(t, err)

	// Test creation
	info, err := keyring.NewInMemory().CreateAccount("xxx", mnemonic, "", "012345678", keyring.CreateHDPath(0, 0).String(), keyring.Secp256k1)
	require.NoError(t, err)
	require.Equal(t, addr, info.GetAddress())
}

func TestGenerateSaveCoinKey(t *testing.T) {
	t.Parallel()
	dir, cleanup := tests.NewTestCaseDir(t)
	t.Cleanup(cleanup)

	kb, err := keyring.NewKeyring(t.Name(), "test", dir, nil)
	require.NoError(t, err)

	addr, mnemonic, err := server.GenerateSaveCoinKey(kb, "keyname", "012345678", false)
	require.NoError(t, err)

	// Test key was actually saved
	info, err := kb.Get("keyname")
	require.NoError(t, err)
	require.Equal(t, addr, info.GetAddress())

	// Test in-memory recovery
	info, err = keyring.NewInMemory().CreateAccount("xxx", mnemonic, "", "012345678", keyring.CreateHDPath(0, 0).String(), keyring.Secp256k1)
	require.NoError(t, err)
	require.Equal(t, addr, info.GetAddress())
}

func TestGenerateSaveCoinKeyOverwriteFlag(t *testing.T) {
	t.Parallel()
	dir, cleanup := tests.NewTestCaseDir(t)
	t.Cleanup(cleanup)

	kb, err := keyring.NewKeyring(t.Name(), "test", dir, nil)
	require.NoError(t, err)

	keyname := "justakey"
	addr1, _, err := server.GenerateSaveCoinKey(kb, keyname, "012345678", false)
	require.NoError(t, err)

	// Test overwrite with overwrite=false
	_, _, err = server.GenerateSaveCoinKey(kb, keyname, "012345678", false)
	require.Error(t, err)

	// Test overwrite with overwrite=true
	addr2, _, err := server.GenerateSaveCoinKey(kb, keyname, "012345678", true)
	require.NoError(t, err)

	require.NotEqual(t, addr1, addr2)
}
