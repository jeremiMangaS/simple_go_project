from youtube_transcript_api import YouTubeTranscriptApi

video_id = 'knVaCNiH-8I'

# access video
ytt_api = YouTubeTranscriptApi()
fetched_transcript = ytt_api.fetch(video_id) 

for snippet in fetched_transcript :
    print(snippet.text)
