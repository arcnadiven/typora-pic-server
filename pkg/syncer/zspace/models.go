package zspace

type CommonFields struct {
	Code    string `json:"code"`
	Ts      int64  `json:"ts"`
	Msg     string `json:"msg"`
	Reason  string `json:"reason"`
	Suggest string `json:"suggest"`
}

func (f *CommonFields) Exist() bool {
	return f.Msg != "文件不存在"
}

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

type ZSpaceFileList struct {
	CommonFields
	Data struct {
		ReqParam struct {
			SortBy string `json:"sort_by"`
			Order  string `json:"order"`
		} `json:"req_param"`
		Info struct {
			Name              string `json:"name"`
			Path              string `json:"path"`
			IsSys             string `json:"is_sys"`
			IsDir             string `json:"is_dir"`
			Type              string `json:"type"`
			CreatedBy         string `json:"created_by"`
			DoubleWrite       string `json:"double_write"`
			DoubleWriteStatus string `json:"double_write_status"`
			DoubleWriteTime   string `json:"double_write_time"`
			Fnum              string `json:"fnum"`
			Vnum              string `json:"vnum"`
			Inum              string `json:"inum"`
			Anum              string `json:"anum"`
			Comnum            string `json:"comnum"`
			Appnum            string `json:"appnum"`
			Docnum            string `json:"docnum"`
			Dirnum            string `json:"dirnum"`
			Tfnum             string `json:"tfnum"`
			Tvnum             string `json:"tvnum"`
			Tinum             string `json:"tinum"`
			Tanum             string `json:"tanum"`
			Tcomnum           string `json:"tcomnum"`
			Tappnum           string `json:"tappnum"`
			Tdocnum           string `json:"tdocnum"`
			Tdirnum           string `json:"tdirnum"`
			Size              string `json:"size"`
			Width             string `json:"width"`
			Height            string `json:"height"`
			Ori               string `json:"ori"`
			Duration          string `json:"duration"`
			Ftype             string `json:"ftype"`
			Longitude         string `json:"longitude"`
			Latitude          string `json:"latitude"`
			FileHash          string `json:"file_hash"`
			UserID            string `json:"user_id"`
			Username          string `json:"username"`
			Nshare            string `json:"nshare"`
			Ext               string `json:"ext"`
			Crtime            string `json:"crtime"`
			Brtime            string `json:"brtime"`
			Dctime            string `json:"dctime"`
			Sttime            string `json:"sttime"`
			ModifyTime        string `json:"modify_time"`
			ChangeTime        string `json:"change_time"`
			AccessTime        string `json:"access_time"`
			Permission        int    `json:"permission"`
			Labels            string `json:"labels"`
			Favorite          bool   `json:"favorite"`
			Encrypted         string `json:"encrypted"`
			EncryptIcon       string `json:"encrypt_icon"`
			OriginalPath      string `json:"original_path"`
			PinTime           int    `json:"pin_time"`
			MediaGroupUUID    string `json:"media_group_uuid"`
			ContentIdentifier string `json:"content_identifier"`
			HasVideo          int    `json:"has_video"`
			FileNote          string `json:"file_note"`
			FileDeleted       string `json:"file_deleted"`
			Mori              string `json:"mori"`
			OffsetTime        string `json:"offset_time"`
		} `json:"info"`
		List  []ZSpaceFile `json:"list"`
		Total string       `json:"total"`
	} `json:"data"`
}

type ZSpaceFile struct {
	Name              string `json:"name"`
	Path              string `json:"path"`
	IsSys             string `json:"is_sys"`
	IsDir             string `json:"is_dir"`
	Type              string `json:"type"`
	CreatedBy         string `json:"created_by"`
	DoubleWrite       string `json:"double_write"`
	DoubleWriteStatus string `json:"double_write_status"`
	DoubleWriteTime   string `json:"double_write_time"`
	Fnum              string `json:"fnum"`
	Vnum              string `json:"vnum"`
	Inum              string `json:"inum"`
	Anum              string `json:"anum"`
	Comnum            string `json:"comnum"`
	Appnum            string `json:"appnum"`
	Docnum            string `json:"docnum"`
	Dirnum            string `json:"dirnum"`
	Tfnum             string `json:"tfnum"`
	Tvnum             string `json:"tvnum"`
	Tinum             string `json:"tinum"`
	Tanum             string `json:"tanum"`
	Tcomnum           string `json:"tcomnum"`
	Tappnum           string `json:"tappnum"`
	Tdocnum           string `json:"tdocnum"`
	Tdirnum           string `json:"tdirnum"`
	Size              string `json:"size"`
	Width             string `json:"width"`
	Height            string `json:"height"`
	Ori               string `json:"ori"`
	Duration          string `json:"duration"`
	Ftype             string `json:"ftype"`
	Longitude         string `json:"longitude"`
	Latitude          string `json:"latitude"`
	FileHash          string `json:"file_hash"`
	UserID            string `json:"user_id"`
	Username          string `json:"username"`
	Nshare            string `json:"nshare"`
	Ext               string `json:"ext"`
	Crtime            string `json:"crtime"`
	Brtime            string `json:"brtime"`
	Dctime            string `json:"dctime"`
	Sttime            string `json:"sttime"`
	ModifyTime        string `json:"modify_time"`
	ChangeTime        string `json:"change_time"`
	AccessTime        string `json:"access_time"`
	Permission        int    `json:"permission"`
	Labels            string `json:"labels"`
	Favorite          bool   `json:"favorite"`
	Encrypted         string `json:"encrypted"`
	EncryptIcon       string `json:"encrypt_icon"`
	OriginalPath      string `json:"original_path"`
	PinTime           int    `json:"pin_time"`
	MediaGroupUUID    string `json:"media_group_uuid"`
	ContentIdentifier string `json:"content_identifier"`
	HasVideo          int    `json:"has_video"`
	FileNote          string `json:"file_note"`
	FileDeleted       string `json:"file_deleted"`
	Mori              string `json:"mori"`
	OffsetTime        string `json:"offset_time"`
}

type ZSpaceFileCreate struct {
	CommonFields
	Data struct {
		Name              string `json:"name"`
		Path              string `json:"path"`
		IsSys             string `json:"is_sys"`
		IsDir             string `json:"is_dir"`
		Type              string `json:"type"`
		CreatedBy         string `json:"created_by"`
		DoubleWrite       string `json:"double_write"`
		DoubleWriteStatus string `json:"double_write_status"`
		DoubleWriteTime   string `json:"double_write_time"`
		Fnum              string `json:"fnum"`
		Vnum              string `json:"vnum"`
		Inum              string `json:"inum"`
		Anum              string `json:"anum"`
		Comnum            string `json:"comnum"`
		Appnum            string `json:"appnum"`
		Docnum            string `json:"docnum"`
		Dirnum            string `json:"dirnum"`
		Tfnum             string `json:"tfnum"`
		Tvnum             string `json:"tvnum"`
		Tinum             string `json:"tinum"`
		Tanum             string `json:"tanum"`
		Tcomnum           string `json:"tcomnum"`
		Tappnum           string `json:"tappnum"`
		Tdocnum           string `json:"tdocnum"`
		Tdirnum           string `json:"tdirnum"`
		Size              string `json:"size"`
		Width             string `json:"width"`
		Height            string `json:"height"`
		Ori               string `json:"ori"`
		Duration          string `json:"duration"`
		Ftype             string `json:"ftype"`
		Longitude         string `json:"longitude"`
		Latitude          string `json:"latitude"`
		FileHash          string `json:"file_hash"`
		UserID            string `json:"user_id"`
		Username          string `json:"username"`
		Nshare            string `json:"nshare"`
		Ext               string `json:"ext"`
		Crtime            string `json:"crtime"`
		Brtime            string `json:"brtime"`
		Dctime            string `json:"dctime"`
		Sttime            string `json:"sttime"`
		ModifyTime        string `json:"modify_time"`
		ChangeTime        string `json:"change_time"`
		AccessTime        string `json:"access_time"`
		Permission        int    `json:"permission"`
		Labels            string `json:"labels"`
		Favorite          bool   `json:"favorite"`
		Encrypted         string `json:"encrypted"`
		EncryptIcon       string `json:"encrypt_icon"`
		OriginalPath      string `json:"original_path"`
		PinTime           int    `json:"pin_time"`
		MediaGroupUUID    string `json:"media_group_uuid"`
		ContentIdentifier string `json:"content_identifier"`
		HasVideo          int    `json:"has_video"`
		FileNote          string `json:"file_note"`
		FileDeleted       string `json:"file_deleted"`
		Mori              string `json:"mori"`
		OffsetTime        string `json:"offset_time"`
	} `json:"data"`
}

type ZSpaceFileRemove struct {
	CommonFields
	Data struct {
		Otype string `json:"otype"`
		Task  struct {
			ID                  string        `json:"id"`
			UserID              string        `json:"user_id"`
			Username            string        `json:"username"`
			Opt                 string        `json:"opt"`
			Src                 []string      `json:"src"`
			LpvSrc              []interface{} `json:"lpv_src"`
			LpvTo               []interface{} `json:"lpv_to"`
			To                  string        `json:"to"`
			DecompressSources   interface{}   `json:"decompress_sources"`
			Pwd                 string        `json:"pwd"`
			Progress            string        `json:"progress"`
			Speed               string        `json:"speed"`
			State               string        `json:"state"`
			Ttype               string        `json:"ttype"`
			Ext                 string        `json:"ext"`
			TotalNum            string        `json:"total_num"`
			TotalSize           string        `json:"total_size"`
			DoneNum             string        `json:"done_num"`
			DoneSize            string        `json:"done_size"`
			FailNum             string        `json:"fail_num"`
			SrcDnum             string        `json:"src_dnum"`
			SrcFnum             string        `json:"src_fnum"`
			Failures            []interface{} `json:"failures"`
			SkipBin             string        `json:"skip_bin"`
			CreatedAt           int           `json:"created_at"`
			UpdatedAt           int           `json:"updated_at"`
			SucceedNum          string        `json:"succeed_num"`
			SrcNum              string        `json:"src_num"`
			MoveNum             string        `json:"move_num"`
			MoveSize            string        `json:"move_size"`
			CopyNum             string        `json:"copy_num"`
			CopySize            string        `json:"copy_size"`
			ErrorSave           string        `json:"error_save"`
			FileHashCheckOption string        `json:"file_hash_check_option"`
		} `json:"task"`
	} `json:"data"`
}

type ZSpaceFileNewDir struct {
	CommonFields
	Data struct {
		Name              string `json:"name"`
		Path              string `json:"path"`
		IsSys             string `json:"is_sys"`
		IsDir             string `json:"is_dir"`
		Type              string `json:"type"`
		CreatedBy         string `json:"created_by"`
		DoubleWrite       string `json:"double_write"`
		DoubleWriteStatus string `json:"double_write_status"`
		DoubleWriteTime   string `json:"double_write_time"`
		Fnum              string `json:"fnum"`
		Vnum              string `json:"vnum"`
		Inum              string `json:"inum"`
		Anum              string `json:"anum"`
		Comnum            string `json:"comnum"`
		Appnum            string `json:"appnum"`
		Docnum            string `json:"docnum"`
		Dirnum            string `json:"dirnum"`
		Tfnum             string `json:"tfnum"`
		Tvnum             string `json:"tvnum"`
		Tinum             string `json:"tinum"`
		Tanum             string `json:"tanum"`
		Tcomnum           string `json:"tcomnum"`
		Tappnum           string `json:"tappnum"`
		Tdocnum           string `json:"tdocnum"`
		Tdirnum           string `json:"tdirnum"`
		Size              string `json:"size"`
		Width             string `json:"width"`
		Height            string `json:"height"`
		Ori               string `json:"ori"`
		Duration          string `json:"duration"`
		Ftype             string `json:"ftype"`
		Longitude         string `json:"longitude"`
		Latitude          string `json:"latitude"`
		FileHash          string `json:"file_hash"`
		UserID            string `json:"user_id"`
		Username          string `json:"username"`
		Nshare            string `json:"nshare"`
		Ext               string `json:"ext"`
		Crtime            string `json:"crtime"`
		Brtime            string `json:"brtime"`
		Dctime            string `json:"dctime"`
		Sttime            string `json:"sttime"`
		ModifyTime        string `json:"modify_time"`
		ChangeTime        string `json:"change_time"`
		AccessTime        string `json:"access_time"`
		Permission        int    `json:"permission"`
		Labels            string `json:"labels"`
		Favorite          bool   `json:"favorite"`
		Encrypted         string `json:"encrypted"`
		EncryptIcon       string `json:"encrypt_icon"`
		OriginalPath      string `json:"original_path"`
		PinTime           int    `json:"pin_time"`
		MediaGroupUUID    string `json:"media_group_uuid"`
		ContentIdentifier string `json:"content_identifier"`
		HasVideo          int    `json:"has_video"`
		FileNote          string `json:"file_note"`
		FileDeleted       string `json:"file_deleted"`
		Mori              string `json:"mori"`
		OffsetTime        string `json:"offset_time"`
	} `json:"data"`
}
