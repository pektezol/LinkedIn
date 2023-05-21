<template>
    <div>
        <!-- left -->
        <div class="card">
            <div class="card-header text-center"
              style="background-image: url(https://image.shutterstock.com/image-vector/seamless-pattern-computer-background-numbers-260nw-1410295883.jpg); height: 50px;">
              <img
                src="https://media-exp1.licdn.com/dms/image/C4D03AQFzRzfBHzBpbA/profile-displayphoto-shrink_400_400/0/1661208519486?e=1669248000&v=beta&t=XFtl1ZvqbPC7otwvvml2B1rOUT7Z-k0dlfLz-DzRq84"
                class="rounded-circle" style="width: 70px; height: 70px; border: 2px solid white;" />
            </div>
            <div class="card-body mt-4" style="font-family: 'Poppins', sans-serif;">
              <p class="card-title header text-center">
                <a href="https://www.linkedin.com/in/alimertkocaman/" target="_blank" style="color: #212529;">Zeynep Bilge
                  ÖZKAN</a>
              </p>
              <div class="text-center" style="margin-bottom: 10px; color: #5f5f5f;">
                Metaverse Developer at Navlungo - Dijital Lojistik Platformu
              </div>
              <hr>
              <div class="card-text">
                <div>
                  <span class="view">Profilinizi kimler görüntüledi </span><span class="float-right view mt-1"
                    style="color: #0a66c2">256</span>
                </div>
                <div>
                  <span class="view">Yayınınız için görüntüleme </span><span class="float-right view mt-1"
                    style="color: #0a66c2">104</span>
                </div>
              </div>
            </div>
            <div class="card-footer ">
              <i class="fas fa-bookmark"></i><span class="mb-4 ml-2">Ögelerim</span>
            </div>
          </div>
       

    </div>
  </template>
  
  <script>
  export default {
    data () {
      return {
        user: {
          data: {
            photoURL: 'https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=200',
            displayName: 'VolkanO6',
            email: 'volkanoztoklu@gmail.com'
          }
        }
      }
    }
  }
  </script>
  