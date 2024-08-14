### 视频流 直播播放 dash

### 生成dash命令

#### 方法一
ffmpeg -re -i input.mkv -map 0:v -map 0:a -c:v libx264 -c:a aac  -f dash output/output.mpd

#### 方法二
ffmpeg -hide_banner -y -threads 0  \
		-i input.mkv  \
        -map v  \
        -c:v:0 libx264 -b:v 800k \
		-map a -c:a:0 aac -b:a 128k  \
		-g 30  \
	 -seg_duration 5 \
		-sc_threshold 0  \
		-b_strategy 0  \
		-use_timeline 0  \
		-use_template 1  \
		-f dash output/output.mpd



#### libx264 和 libx265 的区别

编码标准：

libx264：用于H.264/AVC（Advanced Video Coding），这是目前广泛使用的视频压缩标准，广泛应用于流媒体、蓝光光盘和各种视频平台。
libx265：用于H.265/HEVC（High Efficiency Video Coding），这是H.264的继任者，提供更高的压缩效率，适用于4K视频和高动态范围（HDR）内容。
压缩效率：

libx264：压缩效率较低，但编码速度较快，兼容性好。
libx265：压缩效率更高，可以在相同的画质下使用更低的比特率，但编码速度较慢，编码复杂度更高。
应用场景：

libx264：适用于大多数常见的视频应用，如在线视频、直播、视频会议等。
libx265：适用于需要高压缩效率的应用，如4K视频、存储优化和高质量视频传输。
如何查看macOS下FFmpeg支持的编码器

你可以使用以下命令查看FFmpeg支持的编码器：

ffmpeg -encoders | grep 'libx264\|libx265'


### dash 官网案例

ffmpeg -re -i <input> -map 0 -map 0 -c:a libfdk_aac -c:v libx264 \
-b:v:0 800k -profile:v:0 main \
-b:v:1 300k -s:v:1 320x170 -profile:v:1 baseline -ar:a:1 22050 \
-bf 1 -keyint_min 120 -g 120 -sc_threshold 0 -b_strategy 0 \
-use_timeline 1 -use_template 1 -window_size 5 \
-adaptation_sets "id=0,streams=v id=1,streams=a" \
-f dash /path/to/out.mpd


参数解释 
-re:
以实时速度读取输入。这个选项通常在进行流媒体传输时使用，确保输入数据的读取速度不会超过实际播放速度。

-i <input>:
指定输入文件或流。<input>应该被替换为实际的输入文件路径或URL。
-map 0 -map 0:

映射所有输入流，生成多个输出流。这里-map 0两次表示将输入的所有流（音频、视频）都映射到两个不同的输出中。

-c:a libfdk_aac:

指定音频编码器为libfdk_aac。
-c:v libx264:

指定视频编码器为libx264。
-b:v:0 800k:

为第一个视频流设置比特率为800 kbps。
-profile:v:0 main:

为第一个视频流设置H.264配置文件为main。
-b:v:1 300k:

为第二个视频流设置比特率为300 kbps。
-s:v:1 320x170:

为第二个视频流设置分辨率为320x170。
-profile:v:1 baseline:

为第二个视频流设置H.264配置文件为baseline。
-ar:a:1 22050:

为第二个音频流设置采样率为22050 Hz。
-bf 1:

设置B帧的数量为1（B帧是在视频压缩中用于提高压缩效率的预测帧）。
-keyint_min 120:

设置关键帧之间的最小间隔为120帧。
-g 120:

设置GOP（Group of Pictures）大小为120帧，即每隔120帧插入一个关键帧。
-sc_threshold 0:

设置场景变化的阈值为0，这意味着不会因为场景变化而插入额外的关键帧。
-b_strategy 0:

禁用B帧的策略选择。
-use_timeline 1:

指定DASH使用时间线模式。
-use_template 1:

指定DASH使用模板模式。
-window_size 5:

设置DASH播放窗口的大小为5个片段。
-adaptation_sets "id=0,streams=v id=1,streams=a":

设置DASH自适应集，其中id=0表示视频流，id=1表示音频流。
-f dash:

指定输出格式为DASH。
/path/to/out.mpd:

输出的DASH MPD（媒体呈现描述）文件路径。
总结

这个命令将输入文件或流转换为DASH格式的视频流，生成两个视频流（一个高比特率和一个低比特率）和一个音频流。输出的MPD文件和相关的媒体片段可以用于DASH播放器进行自适应流媒体播放。

