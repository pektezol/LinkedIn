<template>
  <div class="card mt-2 mb-3">
    <div class="card-header">

      <img src="https://miro.medium.com/max/2400/1*dyO41GLMpO6Qdtw8QfjvFw.jpeg" class="postphoto">
      <div class="d-inline-flex flex-column ml-1 align-middle">
        <span class="posttext">
          <a href="https://www.linkedin.com/in/ozanhalis/" target="_blank" style="color: #212529">
            {{ first_name }} {{ last_name }}
          </a>
        </span>
        <span class="profile-desc"> {{ headline }}</span>
        <span class="profile-desc"> {{ date.split("T")[0] }}
          <i class="fas fa-globe-americas fa-sm"></i>
        </span>
      </div>

      <span class="float-right">
        <i class="fas fa-ellipsis-h fa-md"></i>
      </span>
    </div>
    <div>
      <p class="card-text m-3 card-desc">
        {{ content_text }}
      </p>
      <div>
        <img src="https://miro.medium.com/max/1080/1*gdDWFSvHDt57DS2zsh_0Bg.png" class="card-img-top mb-1" alt="..." />
      </div>
      <div class="ml-3 mb-2">


        <span class="profile-desc">
          <img src="https://static-exp1.licdn.com/sc/h/7fx9nkd7mx8avdpqm5hqcbi97" alt="Heart Icon">
          {{ likes }}
        </span>
      </div>
      <div class="card-footer mt-1  text-center">
        <span>
          <button class="ref btn btn-light p-2">
            <i class="far fa-thumbs-up fa-md" style="font-size: 1.2rem">
              <span class="ml-2 mediatext">
                Like
              </span>
            </i>
          </button>
        </span>
        <span>
          <button class="ref btn btn-light p-2 ml-4">
            <i class="far fa-comment fa-md" style="font-size: 1.2rem">
              <span class="ml-2 mediatext">
                Yorum Yap
              </span>
            </i>
          </button>
        </span>
         
      </div>
    </div>
  </div>
</template>
  
<script>
export default {
  name: "post",
  props: ['id', 'first_name', 'last_name', 'headline', 'content_text', 'content_image', 'likes', 'comments', 'date'],
  data() {
    return {

    }
  }
}
</script>
  