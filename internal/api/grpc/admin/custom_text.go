package admin

import (
	"context"

	"golang.org/x/text/language"

	"github.com/caos/zitadel/internal/api/grpc/object"
	text_grpc "github.com/caos/zitadel/internal/api/grpc/text"
	"github.com/caos/zitadel/internal/domain"
	admin_pb "github.com/caos/zitadel/pkg/grpc/admin"
)

func (s *Server) GetDefaultInitMessageText(ctx context.Context, req *admin_pb.GetDefaultInitMessageTextRequest) (*admin_pb.GetDefaultInitMessageTextResponse, error) {
	msg, err := s.query.DefaultMessageTextByTypeAndLanguageFromFileSystem(domain.InitCodeMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetDefaultInitMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) GetCustomInitMessageText(ctx context.Context, req *admin_pb.GetCustomInitMessageTextRequest) (*admin_pb.GetCustomInitMessageTextResponse, error) {
	msg, err := s.query.CustomMessageTextByTypeAndLanguage(ctx, domain.IAMID, domain.InitCodeMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetCustomInitMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) SetDefaultInitMessageText(ctx context.Context, req *admin_pb.SetDefaultInitMessageTextRequest) (*admin_pb.SetDefaultInitMessageTextResponse, error) {
	result, err := s.command.SetDefaultMessageText(ctx, SetInitCustomTextToDomain(req))
	if err != nil {
		return nil, err
	}
	return &admin_pb.SetDefaultInitMessageTextResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) ResetCustomInitMessageTextToDefault(ctx context.Context, req *admin_pb.ResetCustomInitMessageTextToDefaultRequest) (*admin_pb.ResetCustomInitMessageTextToDefaultResponse, error) {
	result, err := s.command.RemoveIAMMessageTexts(ctx, domain.InitCodeMessageType, language.Make(req.Language))
	if err != nil {
		return nil, err
	}
	return &admin_pb.ResetCustomInitMessageTextToDefaultResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) GetDefaultPasswordResetMessageText(ctx context.Context, req *admin_pb.GetDefaultPasswordResetMessageTextRequest) (*admin_pb.GetDefaultPasswordResetMessageTextResponse, error) {
	msg, err := s.query.DefaultMessageTextByTypeAndLanguageFromFileSystem(domain.PasswordResetMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetDefaultPasswordResetMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) GetCustomPasswordResetMessageText(ctx context.Context, req *admin_pb.GetCustomPasswordResetMessageTextRequest) (*admin_pb.GetCustomPasswordResetMessageTextResponse, error) {
	msg, err := s.query.CustomMessageTextByTypeAndLanguage(ctx, domain.IAMID, domain.PasswordResetMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetCustomPasswordResetMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) SetDefaultPasswordResetMessageText(ctx context.Context, req *admin_pb.SetDefaultPasswordResetMessageTextRequest) (*admin_pb.SetDefaultPasswordResetMessageTextResponse, error) {
	result, err := s.command.SetDefaultMessageText(ctx, SetPasswordResetCustomTextToDomain(req))
	if err != nil {
		return nil, err
	}
	return &admin_pb.SetDefaultPasswordResetMessageTextResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) ResetCustomPasswordResetMessageTextToDefault(ctx context.Context, req *admin_pb.ResetCustomPasswordResetMessageTextToDefaultRequest) (*admin_pb.ResetCustomPasswordResetMessageTextToDefaultResponse, error) {
	result, err := s.command.RemoveIAMMessageTexts(ctx, domain.PasswordResetMessageType, language.Make(req.Language))
	if err != nil {
		return nil, err
	}
	return &admin_pb.ResetCustomPasswordResetMessageTextToDefaultResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) GetDefaultVerifyEmailMessageText(ctx context.Context, req *admin_pb.GetDefaultVerifyEmailMessageTextRequest) (*admin_pb.GetDefaultVerifyEmailMessageTextResponse, error) {
	msg, err := s.query.DefaultMessageTextByTypeAndLanguageFromFileSystem(domain.VerifyEmailMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetDefaultVerifyEmailMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) GetCustomVerifyEmailMessageText(ctx context.Context, req *admin_pb.GetCustomVerifyEmailMessageTextRequest) (*admin_pb.GetCustomVerifyEmailMessageTextResponse, error) {
	msg, err := s.query.CustomMessageTextByTypeAndLanguage(ctx, domain.IAMID, domain.VerifyEmailMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetCustomVerifyEmailMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) SetDefaultVerifyEmailMessageText(ctx context.Context, req *admin_pb.SetDefaultVerifyEmailMessageTextRequest) (*admin_pb.SetDefaultVerifyEmailMessageTextResponse, error) {
	result, err := s.command.SetDefaultMessageText(ctx, SetVerifyEmailCustomTextToDomain(req))
	if err != nil {
		return nil, err
	}
	return &admin_pb.SetDefaultVerifyEmailMessageTextResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) ResetCustomVerifyEmailMessageTextToDefault(ctx context.Context, req *admin_pb.ResetCustomVerifyEmailMessageTextToDefaultRequest) (*admin_pb.ResetCustomVerifyEmailMessageTextToDefaultResponse, error) {
	result, err := s.command.RemoveIAMMessageTexts(ctx, domain.VerifyEmailMessageType, language.Make(req.Language))
	if err != nil {
		return nil, err
	}
	return &admin_pb.ResetCustomVerifyEmailMessageTextToDefaultResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) GetDefaultVerifyPhoneMessageText(ctx context.Context, req *admin_pb.GetDefaultVerifyPhoneMessageTextRequest) (*admin_pb.GetDefaultVerifyPhoneMessageTextResponse, error) {
	msg, err := s.query.DefaultMessageTextByTypeAndLanguageFromFileSystem(domain.VerifyPhoneMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetDefaultVerifyPhoneMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) GetCustomVerifyPhoneMessageText(ctx context.Context, req *admin_pb.GetCustomVerifyPhoneMessageTextRequest) (*admin_pb.GetCustomVerifyPhoneMessageTextResponse, error) {
	msg, err := s.query.CustomMessageTextByTypeAndLanguage(ctx, domain.IAMID, domain.VerifyPhoneMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetCustomVerifyPhoneMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) SetDefaultVerifyPhoneMessageText(ctx context.Context, req *admin_pb.SetDefaultVerifyPhoneMessageTextRequest) (*admin_pb.SetDefaultVerifyPhoneMessageTextResponse, error) {
	result, err := s.command.SetDefaultMessageText(ctx, SetVerifyPhoneCustomTextToDomain(req))
	if err != nil {
		return nil, err
	}
	return &admin_pb.SetDefaultVerifyPhoneMessageTextResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) ResetCustomVerifyPhoneMessageTextToDefault(ctx context.Context, req *admin_pb.ResetCustomVerifyPhoneMessageTextToDefaultRequest) (*admin_pb.ResetCustomVerifyPhoneMessageTextToDefaultResponse, error) {
	result, err := s.command.RemoveIAMMessageTexts(ctx, domain.VerifyPhoneMessageType, language.Make(req.Language))
	if err != nil {
		return nil, err
	}
	return &admin_pb.ResetCustomVerifyPhoneMessageTextToDefaultResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) GetDefaultDomainClaimedMessageText(ctx context.Context, req *admin_pb.GetDefaultDomainClaimedMessageTextRequest) (*admin_pb.GetDefaultDomainClaimedMessageTextResponse, error) {
	msg, err := s.query.DefaultMessageTextByTypeAndLanguageFromFileSystem(domain.DomainClaimedMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetDefaultDomainClaimedMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) GetCustomDomainClaimedMessageText(ctx context.Context, req *admin_pb.GetCustomDomainClaimedMessageTextRequest) (*admin_pb.GetCustomDomainClaimedMessageTextResponse, error) {
	msg, err := s.query.CustomMessageTextByTypeAndLanguage(ctx, domain.IAMID, domain.DomainClaimedMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetCustomDomainClaimedMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) SetDefaultDomainClaimedMessageText(ctx context.Context, req *admin_pb.SetDefaultDomainClaimedMessageTextRequest) (*admin_pb.SetDefaultDomainClaimedMessageTextResponse, error) {
	result, err := s.command.SetDefaultMessageText(ctx, SetDomainClaimedCustomTextToDomain(req))
	if err != nil {
		return nil, err
	}
	return &admin_pb.SetDefaultDomainClaimedMessageTextResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) ResetCustomDomainClaimedMessageTextToDefault(ctx context.Context, req *admin_pb.ResetCustomDomainClaimedMessageTextToDefaultRequest) (*admin_pb.ResetCustomDomainClaimedMessageTextToDefaultResponse, error) {
	result, err := s.command.RemoveIAMMessageTexts(ctx, domain.DomainClaimedMessageType, language.Make(req.Language))
	if err != nil {
		return nil, err
	}
	return &admin_pb.ResetCustomDomainClaimedMessageTextToDefaultResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) GetDefaultPasswordlessRegistrationMessageText(ctx context.Context, req *admin_pb.GetDefaultPasswordlessRegistrationMessageTextRequest) (*admin_pb.GetDefaultPasswordlessRegistrationMessageTextResponse, error) {
	msg, err := s.query.DefaultMessageTextByTypeAndLanguageFromFileSystem(domain.PasswordlessRegistrationMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetDefaultPasswordlessRegistrationMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) GetCustomPasswordlessRegistrationMessageText(ctx context.Context, req *admin_pb.GetCustomPasswordlessRegistrationMessageTextRequest) (*admin_pb.GetCustomPasswordlessRegistrationMessageTextResponse, error) {
	msg, err := s.query.CustomMessageTextByTypeAndLanguage(ctx, domain.IAMID, domain.PasswordlessRegistrationMessageType, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetCustomPasswordlessRegistrationMessageTextResponse{
		CustomText: text_grpc.ModelCustomMessageTextToPb(msg),
	}, nil
}

func (s *Server) SetDefaultPasswordlessRegistrationMessageText(ctx context.Context, req *admin_pb.SetDefaultPasswordlessRegistrationMessageTextRequest) (*admin_pb.SetDefaultPasswordlessRegistrationMessageTextResponse, error) {
	result, err := s.command.SetDefaultMessageText(ctx, SetPasswordlessRegistrationCustomTextToDomain(req))
	if err != nil {
		return nil, err
	}
	return &admin_pb.SetDefaultPasswordlessRegistrationMessageTextResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) ResetCustomPasswordlessRegistrationMessageTextToDefault(ctx context.Context, req *admin_pb.ResetCustomPasswordlessRegistrationMessageTextToDefaultRequest) (*admin_pb.ResetCustomPasswordlessRegistrationMessageTextToDefaultResponse, error) {
	result, err := s.command.RemoveIAMMessageTexts(ctx, domain.PasswordlessRegistrationMessageType, language.Make(req.Language))
	if err != nil {
		return nil, err
	}
	return &admin_pb.ResetCustomPasswordlessRegistrationMessageTextToDefaultResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) GetDefaultLoginTexts(ctx context.Context, req *admin_pb.GetDefaultLoginTextsRequest) (*admin_pb.GetDefaultLoginTextsResponse, error) {
	msg, err := s.query.GetDefaultLoginTexts(ctx, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetDefaultLoginTextsResponse{
		CustomText: text_grpc.CustomLoginTextToPb(msg),
	}, nil
}
func (s *Server) GetCustomLoginTexts(ctx context.Context, req *admin_pb.GetCustomLoginTextsRequest) (*admin_pb.GetCustomLoginTextsResponse, error) {
	msg, err := s.query.GetCustomLoginTexts(ctx, domain.IAMID, req.Language)
	if err != nil {
		return nil, err
	}
	return &admin_pb.GetCustomLoginTextsResponse{
		CustomText: text_grpc.CustomLoginTextToPb(msg),
	}, nil
}

func (s *Server) SetCustomLoginText(ctx context.Context, req *admin_pb.SetCustomLoginTextsRequest) (*admin_pb.SetCustomLoginTextsResponse, error) {
	result, err := s.command.SetCustomIAMLoginText(ctx, SetLoginTextToDomain(req))
	if err != nil {
		return nil, err
	}
	return &admin_pb.SetCustomLoginTextsResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}

func (s *Server) ResetCustomLoginTextToDefault(ctx context.Context, req *admin_pb.ResetCustomLoginTextsToDefaultRequest) (*admin_pb.ResetCustomLoginTextsToDefaultResponse, error) {
	result, err := s.command.RemoveCustomIAMLoginTexts(ctx, language.Make(req.Language))
	if err != nil {
		return nil, err
	}
	return &admin_pb.ResetCustomLoginTextsToDefaultResponse{
		Details: object.ChangeToDetailsPb(
			result.Sequence,
			result.EventDate,
			result.ResourceOwner,
		),
	}, nil
}
