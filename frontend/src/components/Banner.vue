<template>
  <div class="banner-wrapper">
    <img src="/img/logotype.png" width="400" />
    <h2>Choose your biggest fear</h2>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Banner extends Vue {}
</script>
<style scoped lang="scss">
.banner-wrapper {
  color: white;
  margin: 0;
}
</style>
