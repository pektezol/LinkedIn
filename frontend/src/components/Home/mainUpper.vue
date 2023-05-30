<template>
    <div class="card">
            <div class="card-body text-center">
              <input type="text" class="form-control content mb-3" placeholder="Bir gönderi başlat" />
              <span><button class="mediatext btn btn-light mr-2">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" width="24"
                    height="24" focusable="false" fill="#70B5F9">
                    <path
                      d="M19 4H5a3 3 0 00-3 3v10a3 3 0 003 3h14a3 3 0 003-3V7a3 3 0 00-3-3zm1 13a1 1 0 01-.29.71L16 14l-2 2-6-6-4 4V7a1 1 0 011-1h14a1 1 0 011 1zm-2-7a2 2 0 11-2-2 2 2 0 012 2z">
                    </path>
                  </svg>
                  <span class="ml-1 mediatext">
                    Fotoğraf
                  </span>
                </button>
              </span>
               
              <span><button class="mediatext btn btn-light mr-2">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" fill=""
                    width="24" height="24" focusable="false">
                    <path d="M21 3v2H3V3zm-6 6h6V7h-6zm0 4h6v-2h-6zm0 4h6v-2h-6zM3 21h18v-2H3zM13 7H3v10h10z">
                    </path>
                  </svg><span class="ml-1 mediatext">
                    Yazı yaz
                  </span>
                </button>
              </span>
            </div>
          </div>
</template>
  
<script>
export default {
    name: "mainUpper",
    data() {
        return {

        }
    },
    methods: {
      encode_file() {
        
      }
    }
}
</script>
  