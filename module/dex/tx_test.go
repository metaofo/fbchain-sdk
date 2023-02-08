package dex

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/metaofo/fbchain-sdk/mocks"
	"github.com/metaofo/fbchain-sdk/module/auth"
	gosdktypes "github.com/metaofo/fbchain-sdk/types"
	"github.com/metaofo/fbchain-sdk/utils"
	"github.com/stretchr/testify/require"
	sdk "github.com/zhengjianfeng1103/fbc/libs/cosmos-sdk/types"
)

func TestDexClient_EditDexOperator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001fibo")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024fibo", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().EditDexOperator(fromInfo, passWd, addr, "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), errors.New("default error"))
	_, err = mockCli.Dex().EditDexOperator(fromInfo, passWd, addr, "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().EditDexOperator(fromInfo, passWd, addr[1:], "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().EditDexOperator(fromInfo, "", addr, "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestDexClient_RegisterDexOperator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001fibo")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024fibo", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().RegisterDexOperator(fromInfo, passWd, addr, "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), errors.New("default error"))
	_, err = mockCli.Dex().RegisterDexOperator(fromInfo, passWd, addr, "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().RegisterDexOperator(fromInfo, passWd, addr[1:], "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().RegisterDexOperator(fromInfo, "", addr, "https://github.com/metaofo/fbchain-sdk", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestDexClient_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001fibo")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024fibo", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().List(fromInfo, passWd, "btc", "fibo", "1.024", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Dex().List(fromInfo, passWd, "", "fibo", "1.024", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().List(fromInfo, passWd, "btc", "", "1.024", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().List(fromInfo, "", "btc", "fibo", "1.024", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	require.Panics(t, func() {
		_, _ = mockCli.Dex().List(fromInfo, passWd, "btc", "fibo", "1.a024", memo,
			accInfo.GetAccountNumber(), accInfo.GetSequence())
	})
}

func TestDexClient_Deposit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001fibo")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024fibo", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().Deposit(fromInfo, passWd, product, "10.24fibo", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Dex().Deposit(fromInfo, passWd, product, "10.24", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().Deposit(fromInfo, passWd, "", "10.24fibo", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().Deposit(fromInfo, "", product, "10.24fibo", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestDexClient_Withdraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001fibo")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024fibo", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().Withdraw(fromInfo, passWd, product, "1.024fibo", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Dex().Withdraw(fromInfo, passWd, "", "1.024fibo", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().Withdraw(fromInfo, "", product, "1.024fibo", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().Withdraw(fromInfo, passWd, product, "1.024", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestDexClient_TransferOwnership(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001fibo")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024fibo", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().TransferOwnership(fromInfo, passWd, product, recAddr, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Dex().TransferOwnership(fromInfo, passWd, "", recAddr, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().TransferOwnership(fromInfo, "", product, recAddr, memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().TransferOwnership(fromInfo, passWd, product, recAddr[1:], memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestDexClient_ConfirmOwnership(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001fibo")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024fibo", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(),
		accInfo.GetSequence()).Return(mocks.DefaultMockSuccessTxResponse(), nil)
	res, err := mockCli.Dex().ConfirmOwnership(fromInfo, passWd, product, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Dex().ConfirmOwnership(fromInfo, passWd, "", memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Dex().ConfirmOwnership(fromInfo, "", product, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}
