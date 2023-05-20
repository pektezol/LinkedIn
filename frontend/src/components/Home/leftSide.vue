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
          <div class="card mt-3">
            <div class="card-header" style="margin-bottom: -22px; font-size: 18px">
              Sayfalarım
            </div>
            <div class="card-body mt-3">
              <div class="card-title mt-2 mb-1">
                <img
                  src="data:image/png;base64,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"
                  style="width: 32px; height: 32px" />
                <div style=" font-size: 15px; font-weight: bold; margin-top: -35px; margin-left: 37px;">
                  Kablosuz Beyin
                </div>
                <div style="font-size: 13px; font-weight: bold; color: gray; margin-left: 38px;">
                  Sayfa Bildirimleri
                  <span class="float-right" style="color: #0a66c2">10</span>
                </div>
              </div>
            </div>
          </div>

          <div class="card mt-3">
            <div class="card-header" style="font-size: 18px">
              En Yeni
            </div>
            <div class="card-body">
              <div class="card-title">
                <div>
                  <i class="fas fa-hashtag" style="font-size: 13px;">
                  </i>
                  <span class="ml-1 font-weight-bold" style="font-size: 13px; color: #5f5f5f">
                    kodluyoruz
                  </span>
                </div>
                <div>
                  <i class="fas fa-hashtag" style="font-size: 13px;">
                  </i>
                  <span class="ml-1 font-weight-bold" style="font-size: 13px; color: #5f5f5f">
                    front-end
                  </span>
                </div>
                <div>
                  <i class="fas fa-hashtag" style="font-size: 13px;">
                  </i>
                  <span class="ml-1 font-weight-bold" style="font-size: 13px; color: #5f5f5f">
                    html
                  </span>
                </div>
                <div>
                  <i class="fas fa-hashtag" style="font-size: 13px;">
                  </i>
                  <span class="ml-1 font-weight-bold" style="font-size: 13px; color: #5f5f5f">
                    css
                  </span>
                </div>
                <div>
                  <i class="fas fa-hashtag" style="font-size: 13px;">
                  </i>
                  <span class="ml-1 font-weight-bold" style="font-size: 13px; color: #5f5f5f">bootstrap
                  </span>
                </div>
                <div class="mt-3 categories">
                  Gruplar
                </div>
                <span class="mt-2 categories">
                  Etkinlikler
                </span>
                <span class="float-right mb-5" style="font-size: 23px; color: #5f5f5f">
                  +
                </span>
                <div class="mt-3 mb-1 categories">
                  Takip edilen hashtagler
                </div>
                <div>
                  <i class="fas fa-hashtag" style="font-size: 13px;">
                  </i>
                  <span class="ml-1 font-weight-bold" style="font-size: 13px; color: #5f5f5f">
                    kodluyoruz
                  </span>
                </div>
                <div>
                  <i class="fas fa-hashtag" style="font-size: 13px;">
                  </i>
                  <span class="ml-1 font-weight-bold" style="font-size: 13px; color: #5f5f5f">
                    front-end
                  </span>
                </div>
              </div>
            </div>
            <div class="card-footer">
              Daha fazla keşfet
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
  