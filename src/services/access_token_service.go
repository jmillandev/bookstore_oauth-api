package services

import (
	"github.com/jmillandev/bookstore_oauth-api/src/domain/access_token"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

type AccessTokenRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type accessTokenService struct {
	repository     AccessTokenRepository
	userRepository UserRepository
}

func NewAccessTokenService(repo AccessTokenRepository, userRepository UserRepository) AccessTokenService {
	return &accessTokenService{
		repository:     repo,
		userRepository: userRepository,
	}
}

func (s *accessTokenService) GetById(accessTokenId string) (*access_token.AccessToken, *rest_errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *accessTokenService) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password
	user, err := s.userRepository.LoginUser(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.repository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *accessTokenService) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
