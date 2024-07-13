package API

type ChannelManager struct {
	subbedChannels []*Channel
}

func (cm *ChannelManager) SubToChannel(channelID string) {
	newChannel := &Channel{ID: channelID}
	cm.subbedChannels = append(cm.subbedChannels, newChannel)
}

func (cm *ChannelManager) PrintMostRecentSubUploads() {

}
