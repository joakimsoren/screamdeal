<template>
  <div class="banner-wrapper">
    <img src="/img/logotype.png" width="400" />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Banner extends Vue {}
</script>
<style scoped lang="scss">
.banner-wrapper {
  font-family: "Lakki Reddy", cursive;
  color: white;
  margin: 0;
}
</style>
