package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"gitlab.33.cn/chain33/chain33/authority/tools/cryptogen/ca"
	"gitlab.33.cn/chain33/chain33/authority/tools/cryptogen/msp"
)

const (
	commonName     = "ca"
	CONFIGFILENAME = "chain33.cryptogen.toml"
	OUTPUTDIR      = "./authdir/crypto"
	ORGNAME        = "Chain33"
)

type Config struct {
	Name []string
}

var (
	cmd = &cobra.Command{
		Use:   "cryptogen [-f configfile] [-o output directory]",
		Short: "chain33 crypto tool for generating key and certificate",
		Run:   generate,
	}
)

func initCfg(path string) *Config {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return &cfg
}

func main() {
	cmd.Flags().StringP("configfile", "f", CONFIGFILENAME, "config file for users")
	cmd.Flags().StringP("outputdir", "o", OUTPUTDIR, "output diraction for key and certificate")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateUsers(baseDir string, users []string, orgName string) {
	fmt.Printf("generateUsers\n")
	fmt.Println(baseDir)

	os.RemoveAll(baseDir)
	caDir := filepath.Join(baseDir, "cacerts")

	signCA, err := ca.NewCA(caDir, commonName)
	if err != nil {
		fmt.Printf("Error generating signCA", err.Error())
		os.Exit(1)
	}

	generateNodes(baseDir, users, signCA, orgName)
}

func generate(cmd *cobra.Command, args []string) {
	configfile, _ := cmd.Flags().GetString("configfile")
	outputdir, _ := cmd.Flags().GetString("outputdir")

	cryptocfg := initCfg(configfile)
	fmt.Println(cryptocfg.Name)

	generateUsers(outputdir, cryptocfg.Name, ORGNAME)
}

func generateNodes(baseDir string, names []string, signCA *ca.CA, orgName string) {
	for _, name := range names {
		userDir := filepath.Join(baseDir, name)
		fileName := fmt.Sprintf("%s@%s", name, orgName)
		err := msp.GenerateLocalMSP(userDir, fileName, signCA)
		if err != nil {
			fmt.Printf("Error generating local MSP")
			os.Exit(1)
		}
	}
}