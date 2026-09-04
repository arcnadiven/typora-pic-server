#!/bin/zsh

cat > ~/Library/LaunchAgents/com.github.typora-pic-server.plist<<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>KeepAlive</key>
	<false/>
	<key>Label</key>
	<string>com.github.typora-pic-server</string>
	<key>ProgramArguments</key>
	<array>
		<string>./bin/typora-pic-server</string>
		<string>--port</string>
		<string>8008</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
	<key>StandardErrorPath</key>
	<string>${HOME}/.go/log/typora-pic-server.log</string>
	<key>StandardOutPath</key>
	<string>${HOME}/.go/log/typora-pic-server.log</string>
	<key>WorkingDirectory</key>
	<string>${HOME}/.go</string>
</dict>
</plist>
EOF

chmod 644 ~/Library/LaunchAgents/com.github.typora-pic-server.plist;
plutil -lint ~/Library/LaunchAgents/com.github.typora-pic-server.plist;
launchctl load -w ~/Library/LaunchAgents/com.github.typora-pic-server.plist;