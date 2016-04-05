package ipfs

import (
	"fmt"
	"io"
	"os"

	core "github.com/ipfs/go-ipfs/core"
	corenet "github.com/ipfs/go-ipfs/core/corenet"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
	peer "gx/ipfs/QmSN2ELGRp4T9kjqiSsSNJRUeR9JKXzQEgwe1HH3tdSGbC/go-libp2p/p2p/peer"
	logging "gx/ipfs/Qmazh5oNUVsDZTs2g59rq8aYQqwpss8tcUWQzor5sCCEuH/go-log"

	"golang.org/x/net/context"
)

func Run(id string) {
	logging.LevelInfo()
	target, err := peer.IDB58Decode(id)
	if err != nil {
		panic(err)
	}

	// Basic ipfsnode setup
	r, err := fsrepo.Open("../.ipfs-repo-client")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := &core.BuildCfg{
		Repo:   r,
		Online: true,
	}

	nd, err := core.NewNode(ctx, cfg)

	if err != nil {
		panic(err)
	}

	fmt.Printf("I am peer %s dialing %s\n", nd.Identity, target)

	con, err := corenet.Dial(nd, target, "/app/lantern")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, con)
}
