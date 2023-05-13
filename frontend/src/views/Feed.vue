<template>
  <div class="feed">
    <FeedNav/>
    <FeedContent/>
  </div>
</template>

<script>
import FeedNav from '@/components/FeedNav.vue'
import FeedContent from '@/components/FeedContent.vue'
export default {
  name: 'Feed',
  components: {
    FeedNav,
    FeedContent
  }
}
</script>
