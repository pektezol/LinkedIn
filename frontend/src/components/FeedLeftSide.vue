<template>
  <section class="feed_content-left ">
    <div class="feed_content-left--identity-card border rounded-tl rounded-tr  text-center pb-10 w-60 bg-white">
      <img class="w-60 h-12 rounded-tl rounded-tr " src="@/assets/images/card-bg.svg" alt="" srcset="">
      <img class="w-20 h-20 relative z-10 -top-8 left-20 w-auto border-2 border-white rounded-full" :src="user.data.photoURL" alt="" srcset="">
      <div class="flex flex-col items-center">
        <h3 class="font-medium">{{user.data.displayName}}</h3>
        <p class="text-xs w-48 text-center">{{user.data.email}}</p>
      </div>
    </div>
    <div class="feed_content-left--identity-card border bg-white py-4 -mt-2 cursor-pointer">
      <div class="hover:bg-gray-100">
        <p class="text-xs font-medium ml-4">Connections</p>
        <h5 class="text-sm font-medium ml-4">Grow your network</h5>
      </div>
    </div>
    <div class="border bg-white py-4 -mt-2 hover:bg-gray-100 cursor-pointer">
      <p class="text-xs font-medium ml-4">Access Exclusive tools and insights</p>
      <h5 class="text-xs flex items-center font-medium ml-4 hover:text-blue-700">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-supported-dps="24x24" class="mercado-match" width="24" height="24" focusable="false">
          <path d="M20 20a3.36 3.36 0 001-2.39V6.38A3.38 3.38 0 0017.62 3H6.38A3.36 3.36 0 004 4z" fill="#f8c77e"></path>
          <path d="M4 4a3.36 3.36 0 00-1 2.38v11.24A3.38 3.38 0 006.38 21h11.24A3.36 3.36 0 0020 20z" fill="#e7a33e"></path>
        </svg>
        Try Premium Free for 1 month
      </h5>
    </div>
    <div class="border bg-white py-4 -mt-2 rounded-b-lg cursor-pointer hover:bg-gray-100">
      <h5 class="text-xs text-gray-900 flex items-center font-medium ml-4">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#6B7280" width="16" height="16" focusable="false">
          <path d="M12 1H4a1 1 0 00-1 1v13.64l5-3.36 5 3.36V2a1 1 0 00-1-1z"></path>
        </svg>
        My items
      </h5>
    </div>
    <div class="flex flex-col items-left mt-6 pt-4 border rounded-tl rounded-tr pb-10 w-60 bg-white sticky top-0">
      <div class="recent ml-3">
        <span class="text-xs text-left">Recent</span>
        <ul class="flex flex-col justify-between mt-4">
          <li class="flex items-center text-xs text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M8.5 7h-1A1.5 1.5 0 006 8.5V14h4V8.5A1.5 1.5 0 008.5 7zM12.75 8h-.5A1.25 1.25 0 0011 9.25V14h3V9.25A1.25 1.25 0 0012.75 8z"></path>
              <circle cx="8" cy="4" r="2"></circle>
              <circle cx="12.5" cy="5.5" r="1.5"></circle>
              <path d="M3.75 8h-.5A1.25 1.25 0 002 9.25V14h3V9.25A1.25 1.25 0 003.75 8z"></path>
              <circle cx="3.5" cy="5.5" r="1.5"></circle>
            </svg>
            Front End Developer Group
          </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M8.5 7h-1A1.5 1.5 0 006 8.5V14h4V8.5A1.5 1.5 0 008.5 7zM12.75 8h-.5A1.25 1.25 0 0011 9.25V14h3V9.25A1.25 1.25 0 0012.75 8z"></path>
              <circle cx="8" cy="4" r="2"></circle>
              <circle cx="12.5" cy="5.5" r="1.5"></circle>
              <path d="M3.75 8h-.5A1.25 1.25 0 002 9.25V14h3V9.25A1.25 1.25 0 003.75 8z"></path>
              <circle cx="3.5" cy="5.5" r="1.5"></circle>
            </svg>
            Software & Technology
          </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M6.65 11h2l-.33 3h2l.33-3H14V9h-3.13l.26-2H14V5h-2.65l.33-3h-2l-.33 3h-2l.33-3h-2l-.33 3H2v2h3.13l-.26 2H2v2h2.65l-.33 3h2zm.48-4h2l-.26 2h-2z"></path>
            </svg>
            careers
            </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M6.65 11h2l-.33 3h2l.33-3H14V9h-3.13l.26-2H14V5h-2.65l.33-3h-2l-.33 3h-2l.33-3h-2l-.33 3H2v2h3.13l-.26 2H2v2h2.65l-.33 3h2zm.48-4h2l-.26 2h-2z"></path>
            </svg>
            management
            </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M6.65 11h2l-.33 3h2l.33-3H14V9h-3.13l.26-2H14V5h-2.65l.33-3h-2l-.33 3h-2l.33-3h-2l-.33 3H2v2h3.13l-.26 2H2v2h2.65l-.33 3h2zm.48-4h2l-.26 2h-2z"></path>
            </svg>
            motivation
          </li>
        </ul>
      </div>
      <div class="groups ml-3">
        <p class="text-xs text-left mt-4 text-blue-700 font-medium">Groups</p>
        <ul class="flex flex-col justify-between mt-2">
          <li class="flex items-center text-xs text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M8.5 7h-1A1.5 1.5 0 006 8.5V14h4V8.5A1.5 1.5 0 008.5 7zM12.75 8h-.5A1.25 1.25 0 0011 9.25V14h3V9.25A1.25 1.25 0 0012.75 8z"></path>
              <circle cx="8" cy="4" r="2"></circle>
              <circle cx="12.5" cy="5.5" r="1.5"></circle>
              <path d="M3.75 8h-.5A1.25 1.25 0 002 9.25V14h3V9.25A1.25 1.25 0 003.75 8z"></path>
              <circle cx="3.5" cy="5.5" r="1.5"></circle>
            </svg>
            Front End Developer Group
          </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M8.5 7h-1A1.5 1.5 0 006 8.5V14h4V8.5A1.5 1.5 0 008.5 7zM12.75 8h-.5A1.25 1.25 0 0011 9.25V14h3V9.25A1.25 1.25 0 0012.75 8z"></path>
              <circle cx="8" cy="4" r="2"></circle>
              <circle cx="12.5" cy="5.5" r="1.5"></circle>
              <path d="M3.75 8h-.5A1.25 1.25 0 002 9.25V14h3V9.25A1.25 1.25 0 003.75 8z"></path>
              <circle cx="3.5" cy="5.5" r="1.5"></circle>
            </svg>
            Software & Technology
          </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M8.5 7h-1A1.5 1.5 0 006 8.5V14h4V8.5A1.5 1.5 0 008.5 7zM12.75 8h-.5A1.25 1.25 0 0011 9.25V14h3V9.25A1.25 1.25 0 0012.75 8z"></path>
              <circle cx="8" cy="4" r="2"></circle>
              <circle cx="12.5" cy="5.5" r="1.5"></circle>
              <path d="M3.75 8h-.5A1.25 1.25 0 002 9.25V14h3V9.25A1.25 1.25 0 003.75 8z"></path>
              <circle cx="3.5" cy="5.5" r="1.5"></circle>
            </svg>
            Developers- Andriod, ios dev....
          </li>
          <li class="text-xs ml-4 text-gray-500 font-medium mb-2">See all</li>
        </ul>
      </div>
      <div class="events ml-3">
        <p class="text-xs text-left mt-4 text-blue-700 font-medium">Events</p>
        <p class="text-xs text-left mt-4 text-blue-700 font-medium">Followed Hastags</p>
        <ul class="flex flex-col justify-between mt-2">
          <li class="flex items-center text-xs text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563"
             class="mercado-match" width="16" height="16" focusable="false">
              <path d="M6.65 11h2l-.33 3h2l.33-3H14V9h-3.13l.26-2H14V5h-2.65l.33-3h-2l-.33 3h-2l.33-3h-2l-.33 3H2v2h3.13l-.26 2H2v2h2.65l-.33 3h2zm.48-4h2l-.26 2h-2z"></path>
            </svg>
            Career
          </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M6.65 11h2l-.33 3h2l.33-3H14V9h-3.13l.26-2H14V5h-2.65l.33-3h-2l-.33 3h-2l.33-3h-2l-.33 3H2v2h3.13l-.26 2H2v2h2.65l-.33 3h2zm.48-4h2l-.26 2h-2z"></path>
            </svg>
            management
          </li>
          <li class="flex items-center text-sm text-gray-500 font-medium mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" data-supported-dps="16x16" fill="#4B5563" class="mercado-match" width="16" height="16" focusable="false">
              <path d="M6.65 11h2l-.33 3h2l.33-3H14V9h-3.13l.26-2H14V5h-2.65l.33-3h-2l-.33 3h-2l.33-3h-2l-.33 3H2v2h3.13l-.26 2H2v2h2.65l-.33 3h2zm.48-4h2l-.26 2h-2z"></path>
            </svg>
            motivation
          </li>
          <li class="flex items-center text-xs ml-4 text-gray-500 font-medium mb-2">Show more
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" preserveAspectRatio="xMinYMin meet" focusable="false" class="lazy-loaded">
              <path d="M8 9l5.93-4L15 6.54l-6.15 4.2a1.5 1.5 0 01-1.69 0L1 6.54 2.07 5z" fill="currentColor"></path>
            </svg>
          </li>
        </ul>
      </div>
    </div>
    <div class="border bg-white py-4 -mt-2 mb-10 rounded-b-lg cursor-pointer hover:bg-gray-100">
      <h5 class="text-sm text-center text-gray-400  font-medium">Discover more</h5>
    </div>
  </section>
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
