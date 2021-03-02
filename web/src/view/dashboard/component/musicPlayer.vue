<template>
  <div>
    <div style="width: 100%">
      <APlayer
               :music="audio[0]"
               :list="audio"
               float

      />
    </div>
  </div>
</template>

<script>
import APlayer from 'vue-aplayer';

export default {
  name: "musicPlayer",
  components: {
    APlayer
  },
  data() {
    return {
      audio: [
        {
          title: '东西（Cover：林俊呈）',
          artist: '纳豆',
          src: 'http://music.163.com/song/media/outer/url?id=1321594530.mp3',
          pic: 'https://p1.music.126.net/5zs7IvmLv7KahY3BFzUmrg==/109951163635241613.jpg?param=300y300', // prettier-ignore
          lrc: 'https://cdn.moefe.org/music/lrc/thing.lrc',
        },
        {
          title: '响喜乱舞（Cover：MARiA）',
          artist: '泠鸢yousa',
          src: 'http://music.163.com/song/media/outer/url?id=1318962459.mp3',
          pic: 'https://p1.music.126.net/AUGVPQ_rVrngDH9ocQrn3Q==/109951163613037822.jpg?param=300y300', // prettier-ignore
          lrc: 'https://cdn.moefe.org/music/lrc/kyoukiranbu.lrc',
        },
        {
          title: '啵唧',
          artist: 'Hanser',
          src: 'http://music.163.com/song/media/outer/url?id=1321424246.mp3',
          pic: 'https://p1.music.126.net/K0-IPcIQ9QFvA0jXTBqoWQ==/109951163636756693.jpg?param=300y300', // prettier-ignore
          lrc: 'https://cdn.moefe.org/music/lrc/kiss.lrc',
        },
      ]
    }
  }
}
</script>

<style scoped>

</style>
