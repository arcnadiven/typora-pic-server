package syncer

type ZSpaceVuex struct {
	State struct {
		App struct {
			DevelpmentMode         bool   `json:"develpmentMode"`
			DevelpmentBase         string `json:"develpmentBase"`
			Device                 string `json:"device"`
			Ua                     string `json:"ua"`
			Version                string `json:"version"`
			Plat                   string `json:"plat"`
			DebugMode              bool   `json:"debugMode"`
			ZoomFactor             int    `json:"zoomFactor"`
			DisableGPUAcceleration int    `json:"disableGPUAcceleration"`
			NoSandBox              int    `json:"noSandBox"`
			UseExecFile            int    `json:"useExecFile"`
			UploadSlice            int    `json:"uploadSlice"`
			UploadProcess          int    `json:"uploadProcess"`
			DownloadByIP           bool   `json:"downloadByIp"`
			ConnectFileServiceMode int    `json:"connectFileServiceMode"`
			ShowLoginPopup         bool   `json:"showLoginPopup"`
			DeviceID               string `json:"deviceId"`
			DownloaderUA           string `json:"downloaderUA"`
			DownloadPath           string `json:"downloadPath"`
			OpenWhenLogin          bool   `json:"openWhenLogin"`
			WindowWidth            int    `json:"windowWidth"`
			WindowHeight           int    `json:"windowHeight"`
			WindowLeft             int    `json:"windowLeft"`
			WindowTop              int    `json:"windowTop"`
			WindowMode             string `json:"windowMode"`
			LocalSyncPort          int    `json:"localSyncPort"`
			Vmsize                 string `json:"vmsize"`
			LocalPort              int    `json:"localPort"`
			ForceColorProfile      string `json:"forceColorProfile"`
		} `json:"app"`
		Nas struct {
			NasID          string        `json:"nasId"`
			NasIP          []interface{} `json:"nasIp"`
			ClientPublicIP string        `json:"clientPublicIp"`
			WebBase        string        `json:"webBase"`
			NasName        string        `json:"nasName"`
			CloudPubKey    string        `json:"cloudPubKey"`
			CloudPubKeyID  string        `json:"cloudPubKeyId"`
			NasPubKey      string        `json:"nasPubKey"`
			Locale         string        `json:"locale"`
			Sign           string        `json:"sign"`
			Color          string        `json:"color"`
			DevicePdt      string        `json:"devicePdt"`
			DeviceMode     string        `json:"deviceMode"`
			DiskNum        int           `json:"diskNum"`
			Series         string        `json:"series"`
		} `json:"nas"`
		Theme struct {
			IsDark            bool   `json:"isDark"`
			LastUserThemePath string `json:"lastUserThemePath"`
		} `json:"theme"`
		User struct {
			Token    string `json:"token"`
			Username string `json:"username"`
			Qcname   string `json:"qcname"`
			KeyName  string `json:"keyName"`
			Avatar   string `json:"avatar"`
			IsMaster int    `json:"isMaster"`
			Userinfo struct {
			} `json:"userinfo"`
			LoginInfo struct {
				LoginMode string `json:"loginMode"`
				LockTime  int    `json:"lockTime"`
				Remember  bool   `json:"remember"`
				KeepLogin bool   `json:"keepLogin"`
				LoginPort int    `json:"loginPort"`
				DeviceIP  string `json:"deviceIp"`
				NasID     string `json:"nasId"`
				LockNow   bool   `json:"lockNow"`
			} `json:"loginInfo"`
			Settings map[string]struct {
				Memo struct {
					Show   bool `json:"show"`
					Width  int  `json:"width"`
					Height int  `json:"height"`
				} `json:"memo"`
				ListenClipboard bool   `json:"listenClipboard"`
				ShowHiddenFile  int    `json:"showHiddenFile"`
				HashCheck       int    `json:"hashCheck"`
				HashCheckOption string `json:"hashCheckOption"`
				ShowSyncIcon    bool   `json:"showSyncIcon"`
				CurrentBg       string `json:"currentBg"`
				DiskCode        string `json:"diskCode"`
			} `json:"settings"`
			HistoryUsers []struct {
				Username   string `json:"username"`
				Qcname     string `json:"qcname"`
				IsMaster   int    `json:"isMaster"`
				Mode       string `json:"mode"`
				IP         string `json:"ip"`
				Token      string `json:"token"`
				NasID      string `json:"nasId"`
				Color      string `json:"color"`
				DeviceMode string `json:"deviceMode"`
				Nickname   string `json:"nickname"`
				TLS        bool   `json:"tls"`
				NasName    string `json:"nasName"`
				Avatar     string `json:"avatar"`
				IsLocal    int    `json:"is_local"`
				T          int64  `json:"t"`
			} `json:"historyUsers"`
			MountHomeDir string `json:"mountHomeDir"`
			MountVolume  string `json:"mountVolume"`
		} `json:"user"`
	} `json:"state"`
}
