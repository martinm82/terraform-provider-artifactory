package local

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/artifactory/resource/repository"
	"github.com/jfrog/terraform-provider-shared/packer"
	"github.com/jfrog/terraform-provider-shared/util"
)

var alpineLocalSchema = util.MergeSchema(
	BaseLocalRepoSchema,
	map[string]*schema.Schema{
		"primary_keypair_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Used to sign index files in Alpine Linux repositories. " +
				"See: https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex",
		},
	},
	repository.RepoLayoutRefSchema("local", "alpine"),
	repository.CompressionFormats,
)

func ResourceArtifactoryLocalAlpineRepository() *schema.Resource {

	type AlpineLocalRepo struct {
		RepositoryBaseParams
		PrimaryKeyPairRef string `hcl:"primary_keypair_ref" json:"primaryKeyPairRef"`
	}

	var unPackLocalAlpineRepository = func(data *schema.ResourceData) (interface{}, string, error) {
		d := &util.ResourceData{ResourceData: data}
		repo := AlpineLocalRepo{
			RepositoryBaseParams: UnpackBaseRepo("local", data, "alpine"),
			PrimaryKeyPairRef:    d.GetString("primary_keypair_ref", false),
		}

		return repo, repo.Id(), nil
	}

	return repository.MkResourceSchema(alpineLocalSchema, packer.Default(alpineLocalSchema), unPackLocalAlpineRepository, func() interface{} {
		return &AlpineLocalRepo{
			RepositoryBaseParams: RepositoryBaseParams{
				PackageType: "alpine",
				Rclass:      "local",
			},
		}
	})
}
