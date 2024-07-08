package API

type Video struct {
	videoID     string
	watchStatus bool
}

type Channel struct {
	username          string
	channelID         string
	uploadsPlaylistID string
}

type ChannelManager struct {
	subbedChannels []*Channel
}

func (cm *ChannelManager) AddChannel(channelID string) {
	newChannel := &Channel{channelID: channelID}
	cm.subbedChannels = append(cm.subbedChannels, newChannel)
}

func (cm *ChannelManager) PrintMostRecentSubUploads() {

}
