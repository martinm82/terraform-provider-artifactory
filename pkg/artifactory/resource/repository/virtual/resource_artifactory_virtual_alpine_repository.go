package virtual

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository"
	"github.com/jfrog/terraform-provider-shared/packer"
	"github.com/jfrog/terraform-provider-shared/util"
)

func ResourceArtifactoryVirtualAlpineRepository() *schema.Resource {

	const packageType = "alpine"

	var alpineVirtualSchema = util.MergeSchema(BaseVirtualRepoSchema, map[string]*schema.Schema{
		"primary_keypair_ref": {
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			Description:      "Primary keypair used to sign artifacts. Default value is empty.",
		},
	}, repository.RepoLayoutRefSchema("virtual", packageType))

	type AlpineVirtualRepositoryParams struct {
		RepositoryBaseParamsWithRetrievalCachePeriodSecs
		PrimaryKeyPairRef string `hcl:"primary_keypair_ref" json:"primaryKeyPairRef"`
	}

	var unpackAlpineVirtualRepository = func(s *schema.ResourceData) (interface{}, string, error) {
		d := &util.ResourceData{ResourceData: s}

		repo := AlpineVirtualRepositoryParams{
			RepositoryBaseParamsWithRetrievalCachePeriodSecs: UnpackBaseVirtRepoWithRetrievalCachePeriodSecs(s, packageType),
			PrimaryKeyPairRef: d.GetString("primary_keypair_ref", false),
		}
		repo.PackageType = packageType
		return &repo, repo.Key, nil
	}

	return repository.MkResourceSchema(alpineVirtualSchema, packer.Default(alpineVirtualSchema), unpackAlpineVirtualRepository, func() interface{} {
		return &AlpineVirtualRepositoryParams{
			RepositoryBaseParamsWithRetrievalCachePeriodSecs: RepositoryBaseParamsWithRetrievalCachePeriodSecs{
				RepositoryBaseParams: RepositoryBaseParams{
					Rclass:      "virtual",
					PackageType: packageType,
				},
			},
		}
	})
}
