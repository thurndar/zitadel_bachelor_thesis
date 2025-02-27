package crtlgitops

import (
	"context"

	"github.com/caos/zitadel/pkg/databases"

	"github.com/caos/orbos/mntr"
	"github.com/caos/orbos/pkg/git"
	"github.com/caos/orbos/pkg/kubernetes"
	orbconfig "github.com/caos/orbos/pkg/orb"
	"github.com/caos/zitadel/operator/database"
	orbdb "github.com/caos/zitadel/operator/database/kinds/orb"
	"github.com/caos/zitadel/operator/zitadel"
	orbz "github.com/caos/zitadel/operator/zitadel/kinds/orb"
)

func DestroyOperator(monitor mntr.Monitor, orbConfigPath string, k8sClient *kubernetes.Client, version *string, gitops bool) error {

	orbConfig, err := orbconfig.ParseOrbConfig(orbConfigPath)
	if err != nil {
		return err
	}

	gitClient := git.New(context.Background(), monitor, "orbos", "orbos@caos.ch")
	if err := gitClient.Configure(orbConfig.URL, []byte(orbConfig.Repokey)); err != nil {
		return err
	}

	dbClient, err := databases.NewConnection(monitor, k8sClient, gitops, orbConfig)
	if err != nil {
		return err
	}

	return zitadel.Destroy(monitor, gitClient, orbz.AdaptFunc("ensure", version, gitops, []string{"zitadel", "iam", "dbconnection"}, dbClient), k8sClient)
}

func DestroyDatabase(monitor mntr.Monitor, orbConfigPath string, k8sClient *kubernetes.Client, version *string, gitops bool) error {

	orbConfig, err := orbconfig.ParseOrbConfig(orbConfigPath)
	if err != nil {
		return err
	}

	gitClient := git.New(context.Background(), monitor, "orbos", "orbos@caos.ch")
	if err := gitClient.Configure(orbConfig.URL, []byte(orbConfig.Repokey)); err != nil {
		return err
	}

	return database.Destroy(monitor, gitClient, orbdb.AdaptFunc("", version, gitops, "operator", "database", "backup"), k8sClient)
}
