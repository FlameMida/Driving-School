<template>
  <div>
    <div style="width: 100%">
      <APlayer
          :music="audio"
          :list="list"
          listMaxHeight="190px"
      />
    </div>
  </div>
</template>

<script>
import APlayer from 'vue-aplayer';
import {getMusic} from "@/api/system";

export default {
  name: "musicPlayer",
  components: {
    APlayer
  },
  created() {
    this.loadMusic()

  },
  data() {
    return {
      audio:
          {
            title: '东西（Cover：林俊呈）',
            artist: '纳豆',
            src: 'http://music.163.com/song/media/outer/url?id=1321594530.mp3',
            pic: 'https://p1.music.126.net/5zs7IvmLv7KahY3BFzUmrg==/109951163635241613.jpg?param=300y300', // prettier-ignore

          },
      list: []
    }
  }, methods: {
    async loadMusic() {
      const {data} = await getMusic();
      this.audio = data.music[0]
      this.list = data.music

    },
  }
}
</script>

<style scoped>

</style>
