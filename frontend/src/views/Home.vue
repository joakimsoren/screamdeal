<template>
  <div class="home">
    <Banner />
  </div>
</template>

<script>
import Banner from "@/components/Banner.vue";
import { Vue, Component } from "vue-property-decorator";

@Component({
  components: {
    Banner
  }
})
export default class Home extends Vue {}
</script>
<style lang="scss" scoped>
.home {
}
</style>
