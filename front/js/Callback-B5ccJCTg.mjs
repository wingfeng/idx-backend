var u=(a,e,o)=>new Promise((l,n)=>{var i=s=>{try{c(o.next(s))}catch(r){n(r)}},t=s=>{try{c(o.throw(s))}catch(r){n(r)}},c=s=>s.done?l(s.value):Promise.resolve(s.value).then(i,t);c((o=o.apply(a,e)).next())});import{_ as f,U as m,o as _,u as p,b as d}from"./bootstrap-t35vpcz_.mjs";import{a6 as k,aF as g,aG as C,aJ as h,m as b,aI as I,aK as x,ao as T}from"../jse/index-index-CqEm-tRS.mjs";const $=k({name:"Callback",mounted(){return u(this,null,function*(){new m(_).signinCallback().then(e=>{console.log("user",e),console.log("response accessToken",e.access_token),p().setAccessToken(e.access_token);const o={userId:e.profile.sub,userName:e.profile.preferred_username,email:e.profile.email,roles:e.profile.roles};d().setUserInfo(o),console.log("stored userInfo",d().userInfo),console.log("accessToken",p().accessToken),this.$router.push("/")})})}}),N=x("h1",null,"OIDC Callback Login Successed",-1);function S(a,e,o,l,n,i){const t=g("RouterLink");return C(),h("div",null,[N,b(t,{to:"/"},{default:I(()=>[T("Go to Home")]),_:1})])}const M=f($,[["render",S]]);export{M as default};
