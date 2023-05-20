<template>
    <!-- Right-Bottom-Mesajlaşma - Start -->
    <nav class="navbar navbar-expand-xl navbar-light bg-white position-sticky fixed-bottom border ml-auto"
      style=" width: 300px; height: 50px; position: -webkit-sticky; border-radius: 5px;  right: 15px;">
      <a class="navbar-brand" href="#">
        <img
          src="https://media-exp1.licdn.com/dms/image/C4D03AQFzRzfBHzBpbA/profile-displayphoto-shrink_400_400/0/1661208519486?e=1669248000&v=beta&t=XFtl1ZvqbPC7otwvvml2B1rOUT7Z-k0dlfLz-DzRq84"
          style="width: 30px; height: 30px; border-radius: 50%;" />
      </a>
      <span style="font-size: 20px; font-family:Arial, Helvetica, sans-serif;">
        <strong>Messages</strong>
      </span>
      <span class="offset-3"><img src="https://static.thenounproject.com/png/729124-200.png" width="15px"
          height="15px"></span>
      <span><i class="fas fa-ellipsis-h fa-sm ml-3 float-right"></i></span>
      <span><i class="fas fa-chevron-up fa-sm ml-3 float-right"></i></span>
    </nav>
</template>

<script>
export default {
    name: "message",
    data() {
        return {

        }
    }
}
</script>