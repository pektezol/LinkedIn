<template>
  <div id="app">
    <h2 class="mobile-message">I'm sure you were expecting a mobile version, but bruhh it's just a clone😶😶</h2>
    <router-view>
    </router-view>
  </div>
</template>

<script>
export default {
}
</script>

<style lang="scss">
*{
  padding: 0px;
  margin: 0px;
  box-sizing: border-box;
}

html, h1, h2, h3, h4, h5,h6{
  padding: 0px;
  margin: 0px;
  font: 100%;
}
.mobile-message{
  display: none;
}

@media (max-width: 600px) {
  #app{
    display: flex;
    justify-content: center;
    align-items: center;
  }
  *{
    background-color: rgb(61, 114, 211);
  }
  .mobile-message{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-content: center;
    color: white;
    margin-top: 300px;
    font-size: 20px;
    text-align: center;
    margin-left: 20px;
    margin-right: 20px;
    font-family:'Courier New', Courier, monospace;
    // background: blueviolet;
    background-size: cover;
    max-height: 100%;
  }
}
</style>
