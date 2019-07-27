#! /bin/bash
cd ~/Workspace/Go-space/src/sis_video_go/web
go install
cd ~/Workspace/Go-space/src/sis_video_go/schedule/
go install
cd ~/Workspace/Go-space/src/sis_video_go/stream_server/
go install
cd ~/Workspace/Go-space/src/sis_video_go/
go install
cp /home/nishiiebie/Workspace/go/bin/web /home/nishiiebie/Workspace/sis_video_ui/
cp /home/nishiiebie/Workspace/go/bin/schedule /home/nishiiebie/Workspace/sis_video_ui/inner
cp /home/nishiiebie/Workspace/go/bin/stream_server /home/nishiiebie/Workspace/sis_video_ui/inner/
cp /home/nishiiebie/Workspace/go/bin/sis_video_go /home/nishiiebie/Workspace/sis_video_ui/
cp -R /home/nishiiebie/Workspace/Go-space/src/sis_video_go/templates /home/nishiiebie/Workspace/sis_video_ui/
