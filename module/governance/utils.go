package governance

import (
	"encoding/json"
	"io/ioutil"

	"github.com/metaofo/fbchain-sdk/module/governance/types"
	"github.com/metaofo/fbchain-sdk/utils"
)

func parseProposalFromFile(path string) (proposal types.ProposalJSON, err error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if err = json.Unmarshal(contents, &proposal); err != nil {
		return proposal, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
