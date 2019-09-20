package status

type StatusMsg struct {
	ViewID    int64  `json:"viewID"`
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"clientID"`
	NodeID    string `json:"nodeID"`
	NodeTPS   string `json:"nodeTPS"`
	Trans	  string `json:"trans"`
	ValueA    string `json:"valueA"`
	ValueB    string `json:"valueB"`
	ValueC    string `json:"valueC"`
}