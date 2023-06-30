package lndpay

import (
	"context"

	"github.com/lightninglabs/lndclient"
	"github.com/lightningnetwork/lnd/lnrpc"
)

type Service struct {
	context     context.Context
	lnrpcClient lnrpc.LightningClient
}

func NewInstance(lndHost string, tlsPath string, macarroonDir string, networkType string) (*Service, error) {
	lnrpcClient, err := lndclient.NewBasicClient(lndHost, tlsPath, macarroonDir, networkType, lndclient.MacFilename("admin.macaroon"))
	if err != nil {
		return nil, err
	}

	service := &Service{
		context:     context.Background(),
		lnrpcClient: lnrpcClient,
	}
	return service, nil
}

// Get Infomation
func (s *Service) GetInfo() (*lnrpc.GetInfoResponse, error) {
	infoRequest := &lnrpc.GetInfoRequest{}
	return s.lnrpcClient.GetInfo(s.context, infoRequest)
}

func (s *Service) GetNetworkInfo() (*lnrpc.NetworkInfo, error) {
	infoRequest := &lnrpc.NetworkInfoRequest{}
	return s.lnrpcClient.GetNetworkInfo(s.context, infoRequest)
}

func (s *Service) GetTransactions() (*lnrpc.TransactionDetails, error) {
	infoRequest := &lnrpc.GetTransactionsRequest{}
	return s.lnrpcClient.GetTransactions(s.context, infoRequest)
}

func (s *Service) GetChannelBalance() (*lnrpc.ChannelBalanceResponse, error) {
	infoRequest := &lnrpc.ChannelBalanceRequest{}
	return s.lnrpcClient.ChannelBalance(s.context, infoRequest)
}

// List channels, payments made and Invoices
func (s *Service) ListChannels() (*lnrpc.ListChannelsResponse, error) {
	infoRequest := &lnrpc.ListChannelsRequest{}
	return s.lnrpcClient.ListChannels(s.context, infoRequest)
}

func (s *Service) ListPayments() (*lnrpc.ListPaymentsResponse, error) {
	infoRequest := &lnrpc.ListPaymentsRequest{}
	return s.lnrpcClient.ListPayments(s.context, infoRequest)
}

func (s *Service) ListInvoices() (*lnrpc.ListInvoiceResponse, error) {
	infoRequest := &lnrpc.ListInvoiceRequest{}
	return s.lnrpcClient.ListInvoices(s.context, infoRequest)
}

// Add Invoice
func (s *Service) AddInvoice() (*lnrpc.AddInvoiceResponse, error) {
	infoRequest := &lnrpc.Invoice{}
	return s.lnrpcClient.AddInvoice(s.context, infoRequest)
}
