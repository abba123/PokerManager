webpackJsonp([1],{"8Bw/":function(t,e){},"9M+g":function(t,e){},Jmt5:function(t,e){},LveR:function(t,e){},NHnr:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=a("7+uW"),n={render:function(){var t=this.$createElement,e=this._self._c||t;return e("b-card",{attrs:{title:"Card Title","no-body":"",id:"app"}},[e("b-card-header",{attrs:{"header-tag":"nav"}},[e("b-nav",{attrs:{"card-header":"",tabs:""}},[e("b-nav-item",{attrs:{to:"/",exact:"","exact-active-class":"active"}},[this._v("Home")]),this._v(" "),e("b-nav-item",{attrs:{to:"/getwinrate",exact:"","exact-active-class":"active"}},[this._v("getwinrate")]),this._v(" "),e("b-nav-item",{attrs:{to:"/handmanager",exact:"","exact-active-class":"active"}},[this._v("handmanager")])],1)],1),this._v(" "),e("b-card-body",[e("router-view")],1)],1)},staticRenderFns:[]};var i=a("VU/8")({name:"App",data:function(){return{msg:"Welcome to PokerManager"}}},n,!1,function(t){a("8Bw/")},null,null).exports,o=a("/ocq"),l={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"hello"},[a("b-button",{on:{click:t.getWinRate}},[t._v("Calculate Win Rate")]),t._v(" "),a("table",{staticClass:"table"},[t._m(0),t._v(" "),a("tr",[a("td",{attrs:{rowspan:"2"}},[t._v("Player1")]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player1Card1Num,expression:"player1Card1Num"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player1Card1Num=e.target.multiple?a:a[0]}}},t._l(t.nums,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player1Card1Suit,expression:"player1Card1Suit"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player1Card1Suit=e.target.multiple?a:a[0]}}},t._l(t.suits,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player1Card2Num,expression:"player1Card2Num"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player1Card2Num=e.target.multiple?a:a[0]}}},t._l(t.nums,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player1Card2Suit,expression:"player1Card2Suit"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player1Card2Suit=e.target.multiple?a:a[0]}}},t._l(t.suits,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",{attrs:{id:"result",colspan:"2"}},[t._v("\n          "+t._s(t.winRate1)+"\n        ")])]),t._v(" "),a("tr",[a("td",{attrs:{colspan:"2"}},[a("img",{attrs:{src:"../../static/images/"+t.player1Card1Num+t.player1Card1Suit+".png"}})]),t._v(" "),a("td",{attrs:{colspan:"2"}},[a("img",{attrs:{src:"../../static/images/"+t.player1Card2Num+t.player1Card2Suit+".png"}})])]),t._v(" "),a("tr",[a("td",{attrs:{rowspan:"2"}},[t._v("Player2")]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player2Card1Num,expression:"player2Card1Num"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player2Card1Num=e.target.multiple?a:a[0]}}},t._l(t.nums,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player2Card1Suit,expression:"player2Card1Suit"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player2Card1Suit=e.target.multiple?a:a[0]}}},t._l(t.suits,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player2Card2Num,expression:"player2Card2Num"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player2Card2Num=e.target.multiple?a:a[0]}}},t._l(t.nums,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",[a("select",{directives:[{name:"model",rawName:"v-model",value:t.player2Card2Suit,expression:"player2Card2Suit"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.player2Card2Suit=e.target.multiple?a:a[0]}}},t._l(t.suits,function(e){return a("option",{key:e},[t._v("\n              "+t._s(e)+"\n            ")])}),0)]),t._v(" "),a("td",{attrs:{id:"result",colspan:"2"}},[t._v("\n          "+t._s(t.winRate2)+"\n        ")])]),t._v(" "),a("tr",[a("td",{attrs:{colspan:"2"}},[a("img",{attrs:{src:"../../static/images/"+t.player2Card1Num+t.player2Card1Suit+".png"}})]),t._v(" "),a("td",{attrs:{colspan:"2"}},[a("img",{attrs:{src:"../../static/images/"+t.player2Card2Num+t.player2Card2Suit+".png"}})])])])],1)},staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("thead",[e("th"),this._v(" "),e("th",{attrs:{colspan:"2"}},[this._v("Card1")]),this._v(" "),e("th",{attrs:{colspan:"2"}},[this._v("Crad2")]),this._v(" "),e("th",{attrs:{colspan:"2"}},[this._v("Winrate")])])}]};var s=a("VU/8")({name:"getwinrate",data:function(){return{counter:1,msg:"Welcome to PokerManager",suits:["s","h","d","c"],nums:["1","2","3","4","5","6","7","8","9","10","11","12","13"],player1Card1Num:"",player1Card1Suit:"",player1Card2Num:"",player1Card2Suit:"",player2Card1Num:"",player2Card1Suit:"",player2Card2Num:"",player2Card2Suit:"",player1Card1:"",player1Card2:"",player2Card1:"",player2Card2:"",winRate1:0,winRate2:0}},methods:{getWinRate:function(){var t=this;this.$http.get("http://"+this.$root.backIP+"/getwinrate/",{params:{name1:"player1",name2:"player2",p1Card1Num:this.player1Card1Num,p1Card1Suit:this.player1Card1Suit,p1Card2Num:this.player1Card2Num,p1Card2Suit:this.player1Card2Suit,p2Card1Num:this.player2Card1Num,p2Card1Suit:this.player2Card1Suit,p2Card2Num:this.player2Card2Num,p2Card2Suit:this.player2Card2Suit}}).then(function(e){t.winRate1=e.data.player1,t.winRate2=e.data.player2})}}},l,!1,function(t){a("TznE")},"data-v-9438d2fa",null).exports,u={name:"handmanager",data:function(){return{msg:"Welcome to PokerManager",formData:new FormData,table:[],imgsrc:"../../static/images/",num:1,gain:"all",seat:"all"}},methods:{fileChange:function(t){this.formData.append("file",t.target.files[0])},upload:function(){var t=this;this.$http.put("http://"+this.$root.backIP+"/hand/",this.formData).then(function(e){t.num=10,t.gethand()})},gethand:function(){var t=this;this.$http.get("http://"+this.$root.backIP+"/hand/",{params:{num:this.num,gain:this.gain,seat:this.seat}}).then(function(e){t.table=e.data,console.log(t.table)})}}},p={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"hello"},[a("input",{attrs:{type:"file"},on:{change:t.fileChange}}),t._v(" "),a("b-button",{on:{click:t.upload}},[t._v("upload")]),t._v(" "),a("b-form",[t._v("\n    筆數:\n    "),a("select",{directives:[{name:"model",rawName:"v-model",value:t.num,expression:"num"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.num=e.target.multiple?a:a[0]}}},[a("option",[t._v("1")]),t._v(" "),a("option",[t._v("10")]),t._v(" "),a("option",[t._v("100")])]),t._v("\n    Win/Lose:\n    "),a("select",{directives:[{name:"model",rawName:"v-model",value:t.gain,expression:"gain"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.gain=e.target.multiple?a:a[0]}}},[a("option",[t._v("all")]),t._v(" "),a("option",[t._v(">1.0")]),t._v(" "),a("option",[t._v(">0.5")]),t._v(" "),a("option",[t._v(">0.0")]),t._v(" "),a("option",[t._v(">-0.5")]),t._v(" "),a("option",[t._v(">-1.0")])]),t._v("\n    位置:\n    "),a("select",{directives:[{name:"model",rawName:"v-model",value:t.seat,expression:"seat"}],on:{change:function(e){var a=Array.prototype.filter.call(e.target.options,function(t){return t.selected}).map(function(t){return"_value"in t?t._value:t.value});t.seat=e.target.multiple?a:a[0]}}},[a("option",[t._v("all")]),t._v(" "),a("option",[t._v("HJ")]),t._v(" "),a("option",[t._v("LJ")]),t._v(" "),a("option",[t._v("CO")]),t._v(" "),a("option",[t._v("BTN")]),t._v(" "),a("option",[t._v("SB")]),t._v(" "),a("option",[t._v("BB")])]),t._v(" "),a("b-button",{staticClass:"btn",on:{click:t.gethand}},[t._v("搜尋")])],1),t._v(" "),a("table",[t._m(0),t._v(" "),t._l(t.table,function(e){return a("tr",[a("td",[t._v(t._s(e.Time))]),t._v(" "),a("td",[t._v(t._s(e.Player[0].Name))]),t._v(" "),a("td",[t._v(t._s(e.Player[0].Seat))]),t._v(" "),a("td",[t._v(t._s(e.Player[0].Gain))]),t._v(" "),a("img",{attrs:{src:t.imgsrc+e.Player[0].Card[0].Num+e.Player[0].Card[0].Suit+".png"}}),t._v(" "),a("img",{attrs:{src:t.imgsrc+e.Player[0].Card[1].Num+e.Player[0].Card[1].Suit+".png"}}),t._v(" "),a("td",[t._v(t._s(e.Player[0].Action.Preflop))]),t._v(" "),a("td",[t._v(t._s(e.Player[0].Action.Flop))]),t._v(" "),a("td",[t._v(t._s(e.Player[0].Action.Turn))]),t._v(" "),a("td",[t._v(t._s(e.Player[0].Action.River))])])})],2)],1)},staticRenderFns:[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("tr",[a("td",[t._v("Time")]),t._v(" "),a("td",[t._v("Player")]),t._v(" "),a("td",[t._v("Seat")]),t._v(" "),a("td",[t._v("Gain")]),t._v(" "),a("td",[t._v("Card")]),t._v(" "),a("td",[t._v("Preflop")]),t._v(" "),a("td",[t._v("Flop")]),t._v(" "),a("td",[t._v("Turn")]),t._v(" "),a("td",[t._v("River")])])}]};var c=a("VU/8")(u,p,!1,function(t){a("pCeK")},"data-v-7369c9ec",null).exports,v={render:function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"hello",attrs:{size:"10px"}},[this.$root.token?t._e():a("b-form",{staticStyle:{width:"30%",margin:"auto"}},[this.$root.token?t._e():a("b-form-group",{attrs:{label:"username"}},[a("b-form-input",{attrs:{required:""},model:{value:t.username,callback:function(e){t.username=e},expression:"username"}})],1),t._v(" "),this.$root.token?t._e():a("b-form-group",{attrs:{label:"password"}},[a("b-form-input",{attrs:{required:""},model:{value:t.password,callback:function(e){t.password=e},expression:"password"}})],1),t._v(" "),a("b-button",{staticClass:"btn",attrs:{variant:"primary"},on:{click:t.login}},[t._v("登入")]),t._v(" "),a("b-button",{staticClass:"btn",attrs:{variant:"primary"},on:{click:t.register}},[t._v("註冊")])],1),t._v(" "),this.$root.token?a("div",[a("b-button",{staticClass:"btn",attrs:{variant:"danger"},on:{click:t.logout}},[t._v("登出")])],1):t._e()],1)},staticRenderFns:[]};var d=a("VU/8")({name:"login",data:function(){return{msg:"Welcome to PokerManager",username:"test",password:"test"}},methods:{login:function(){var t=this;this.$http.post("http://"+this.$root.backIP+"/",{username:this.username,password:this.password}).then(function(e){e.data&&(t.$root.token=e.data,t.$http.defaults.headers.common.Authorization=t.$root.token)})},register:function(){this.$http.put("http://"+this.$root.backIP+"/",{username:this.username,password:this.password})},logout:function(){this.$http.delete("http://"+this.$root.backIP+"/"),this.$root.token="",this.$http.defaults.headers.common.Authorization=this.$root.token}}},v,!1,function(t){a("LveR")},"data-v-ff2f50a0",null).exports;r.default.use(o.a);var m=new o.a({routes:[{path:"/",name:"login",component:d,meta:{title:"Entrance"}},{path:"/handmanager",name:"handmanager",component:c,beforeEnter:function(t,e,a){""==r.default.prototype.$http.defaults.headers.common.Authorization&&a({name:"login"}),a()}},{path:"/getwinrate",name:"getwinrate",component:s,beforeEnter:function(t,e,a){""==r.default.prototype.$http.defaults.headers.common.Authorization&&a({name:"login"}),a()}}]}),_=a("Tqaz"),h=a("mtWM"),f=a.n(h);a("Jmt5"),a("9M+g");r.default.config.productionTip=!1,r.default.use(_.a),r.default.prototype.$http=f.a,r.default.prototype.$http.defaults.headers.common.Authorization="",new r.default({el:"#app",router:m,components:{App:i},template:"<App/>",data:function(){return{token:"",backIP:"3.133.150.55"}}})},TznE:function(t,e){},pCeK:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.7579668e2ed7304488a0.js.map