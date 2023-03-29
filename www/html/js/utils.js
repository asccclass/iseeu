// 測試資料是否為JSON格式
const isJSON = (str) => {
   try {
      JSON.parse(str);
   } catch (e){
      return false;
   }
   return true;
}

// member joined
let handleUserJoined = async(response) => {
   offer = JSON.parse(response.msg);
   console.log("someone joined:", offer);
};

// 從URL取得room ID
let getRoomID = () => {
   const queryString = window.location.search;
   const urlParams = new URLSearchParams(queryString);
   return urlParams.get('roomID');
};

// 將回傳的資料轉回 string
let TransObj = (s)=> {
   return JSON.parse(atob(s));
};

// 物件->字串->編碼
let TransBase64 = (s)=> {
   return btoa(JSON.stringify(s));
}

// 傳送資料到server去
let sendMessageToPeer = async(action, str, memberID) => {
   envelope.action = action;
   envelope.to = "" + memberID.toString();
   envelope.msg = TransBase64(str);
   var msg = JSON.stringify(envelope);
   await fetch("/iseeu/sendsocketmessageinstring", {
      method: 'POST',
      // headers: { 'Accept': 'application/json', 'Content-Type': 'application/json' },
      body: msg
   });
};
