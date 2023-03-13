package services

import (
	"context"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/types"
)

func TestEntitiesNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.New(context.TODO(), &types.CreateAnEntityParameters{Structure: increase.P(types.CreateAnEntityParametersStructureCorporation), Corporation: increase.P(types.CreateAnEntityParametersCorporation{Name: increase.P("National Phonograph Company"), Website: increase.P("https://example.com"), TaxIdentifier: increase.P("602214076"), IncorporationState: increase.P("NY"), Address: increase.P(types.CreateAnEntityParametersCorporationAddress{Line1: increase.P("33 Liberty Street"), Line2: increase.P("x"), City: increase.P("New York"), State: increase.P("NY"), Zip: increase.P("10045")}), BeneficialOwners: increase.P([]types.CreateAnEntityParametersCorporationBeneficialOwners{{Individual: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividual{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification{Method: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})}), CompanyTitle: increase.P("x"), Prong: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersProngOwnership)}, {Individual: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividual{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification{Method: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})}), CompanyTitle: increase.P("x"), Prong: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersProngOwnership)}, {Individual: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividual{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentification{Method: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersIndividualIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})}), CompanyTitle: increase.P("x"), Prong: increase.P(types.CreateAnEntityParametersCorporationBeneficialOwnersProngOwnership)}})}), NaturalPerson: increase.P(types.CreateAnEntityParametersNaturalPerson{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersNaturalPersonAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersNaturalPersonIdentification{Method: increase.P(types.CreateAnEntityParametersNaturalPersonIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersNaturalPersonIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersNaturalPersonIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersNaturalPersonIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})}), Joint: increase.P(types.CreateAnEntityParametersJoint{Name: increase.P("x"), Individuals: increase.P([]types.CreateAnEntityParametersJointIndividuals{{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersJointIndividualsAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersJointIndividualsIdentification{Method: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})}, {Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersJointIndividualsAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersJointIndividualsIdentification{Method: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})}, {Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersJointIndividualsAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersJointIndividualsIdentification{Method: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersJointIndividualsIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})}})}), Trust: increase.P(types.CreateAnEntityParametersTrust{Name: increase.P("x"), Category: increase.P(types.CreateAnEntityParametersTrustCategoryRevocable), TaxIdentifier: increase.P("x"), FormationState: increase.P("x"), Address: increase.P(types.CreateAnEntityParametersTrustAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), FormationDocumentFileID: increase.P("string"), Trustees: increase.P([]types.CreateAnEntityParametersTrustTrustees{{Structure: increase.P(types.CreateAnEntityParametersTrustTrusteesStructureIndividual), Individual: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividual{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentification{Method: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})})}, {Structure: increase.P(types.CreateAnEntityParametersTrustTrusteesStructureIndividual), Individual: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividual{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentification{Method: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})})}, {Structure: increase.P(types.CreateAnEntityParametersTrustTrusteesStructureIndividual), Individual: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividual{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentification{Method: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersTrustTrusteesIndividualIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})})}}), Grantor: increase.P(types.CreateAnEntityParametersTrustGrantor{Name: increase.P("x"), DateOfBirth: increase.P("2019-12-27"), Address: increase.P(types.CreateAnEntityParametersTrustGrantorAddress{Line1: increase.P("x"), Line2: increase.P("x"), City: increase.P("x"), State: increase.P("x"), Zip: increase.P("x")}), ConfirmedNoUsTaxID: increase.P(true), Identification: increase.P(types.CreateAnEntityParametersTrustGrantorIdentification{Method: increase.P(types.CreateAnEntityParametersTrustGrantorIdentificationMethodSocialSecurityNumber), Number: increase.P("xxxx"), Passport: increase.P(types.CreateAnEntityParametersTrustGrantorIdentificationPassport{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), Country: increase.P("x")}), DriversLicense: increase.P(types.CreateAnEntityParametersTrustGrantorIdentificationDriversLicense{FileID: increase.P("string"), ExpirationDate: increase.P("2019-12-27"), State: increase.P("x")}), Other: increase.P(types.CreateAnEntityParametersTrustGrantorIdentificationOther{Country: increase.P("x"), Description: increase.P("x"), ExpirationDate: increase.P("2019-12-27"), FileID: increase.P("string")})})})}), Description: increase.P("x"), Relationship: increase.P(types.CreateAnEntityParametersRelationshipAffiliated), SupplementalDocuments: increase.P([]types.CreateAnEntityParametersSupplementalDocuments{{FileID: increase.P("string")}, {FileID: increase.P("string")}, {FileID: increase.P("string")}})})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEntitiesGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.Get(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
	)
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestEntitiesListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.List(context.TODO(), &types.EntityListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0))})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
