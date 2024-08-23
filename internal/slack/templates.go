package slack

func GetDealInfoBlocks() (Blocks, error) {
	message := "Good job"
	blocks := []SectionBlock{
		{Type: "section", Text: TextBlock{Type: "mrkdwn", Text: message}},
	}

	return NewBlocksMap(blocks)
}
