package core

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/textileio/textile-go/pb"
	"github.com/textileio/textile-go/repo"
	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
)

// AddComment adds an outgoing comment block
func (t *Thread) AddComment(target string, body string) (mh.Multihash, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	// build block
	msg := &pb.ThreadComment{
		Target: target,
		Body:   body,
	}

	// commit to ipfs
	res, err := t.commitBlock(msg, pb.ThreadBlock_COMMENT, nil)
	if err != nil {
		return nil, err
	}

	// index it locally
	if err := t.indexBlock(res, repo.CommentBlock, target, body); err != nil {
		return nil, err
	}

	// update head
	if err := t.updateHead(res.hash); err != nil {
		return nil, err
	}

	// post it
	if err := t.post(res, t.Peers()); err != nil {
		return nil, err
	}

	log.Debugf("added COMMENT to %s: %s", t.Id, res.hash.B58String())

	// all done
	return res.hash, nil
}

// handleCommentBlock handles an incoming comment block
func (t *Thread) handleCommentBlock(hash mh.Multihash, block *pb.ThreadBlock) (*pb.ThreadComment, error) {
	msg := new(pb.ThreadComment)
	if err := ptypes.UnmarshalAny(block.Payload, msg); err != nil {
		return nil, err
	}

	// index it locally
	if err := t.indexBlock(&commitResult{
		hash:   hash,
		header: block.Header,
	}, repo.CommentBlock, msg.Target, msg.Body); err != nil {
		return nil, err
	}
	return msg, nil
}