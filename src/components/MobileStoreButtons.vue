<template>
  <div class="store-buttons">
    <a href="https://apps.apple.com/ca/app/ref-buddy/id6444165901" target="_blank">
      <img src="@/assets/app-store-badge.png" alt="Available on the App Store" />
    </a>
    <a href="https://play.google.com/store/apps/details?id=com.duckwxrth.refbuddy"  target="_blank">
      <img src="@/assets/google-play-badge.png" alt="Available on Google Play" />
    </a>
  </div>
</template>

<script>
export default {
  name: 'MobileStores',
  mounted() {
    console.log('MobileStores component mounted.');
  },
};
</script>

<style>
.store-buttons {
  display: flex;
  justify-content: center;
  align-items: center;
}

.store-buttons a {
  margin: 0 10px;
  padding-bottom: 30px;
}

.store-buttons img {
  height: 50px;
  border-radius: 10px;
}
</style>