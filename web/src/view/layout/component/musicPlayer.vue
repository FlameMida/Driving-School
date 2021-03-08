<template>
  <div>
    <APlayer
        :audio="audio"
        :fixed="true"
        :listFolded="true"
        :mini="false"
    />
  </div>
</template>

<script>
import {APlayer} from '@moefe/vue-aplayer';
import {getMusic} from "@/api/system";

export default {
  name: "musicPlayer",
  components: {
    APlayer
  },
  mounted() {
    this.loadMusic()
  },
  data() {
    return {
      audio:
          [
            {
              name: '东西（Cover：林俊呈）',
              artist: '纳豆',
              url: 'http://music.163.com/song/media/outer/url?id=1321594530.mp3',
              cover: 'https://p1.music.126.net/5zs7IvmLv7KahY3BFzUmrg==/109951163635241613.jpg?param=300y300', // prettier-ignore

            },
          ]
      // list: []
    }
  }, methods: {
    async loadMusic() {
      const {data} = await getMusic();
      this.audio = data.music
      // this.list = data.music

    },
  }
}
</script>

<style scoped>

</style>
