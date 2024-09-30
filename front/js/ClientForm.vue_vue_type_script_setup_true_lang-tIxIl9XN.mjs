import{a6 as w,J as s,P as G,aF as u,aG as N,aH as P,aI as n,m as a,aK as L,aW as D,ao as c}from"../jse/index-index-CqEm-tRS.mjs";const E=new Map([["authorization_code","Authorization Code"],["client_credentials","Client Credentials"],["implicit","Implicit"],["password","Password"],["refresh_token","Refresh Token"],["urn:ietf:params:oauth:grant-type:device_code","Device Code"]]),F=w({__name:"ClientForm",props:{client:{}},setup(C,{expose:U}){const _=C,e=s(()=>_.client),p=G(),k=[{label:"openid",value:"openid"},{label:"profile",value:"profile"},{label:"email",value:"email"},{label:"address",value:"address"},{label:"phone",value:"phone"}],v=s({get(){return e.value.Scopes.split(" ")},set(o){e.value.Scopes=o.join(" ")}}),y=()=>p.value.validate().then(()=>!0).catch(o=>(console.log("error submit:",o),!1)),R=()=>{p.value.resetFields()},h=s(()=>{const o=[];return E.forEach((l,r)=>{o.push({label:l,value:r})}),o}),m=s({get(){return e.value.GrantTypes.split(" ")},set(o){e.value.GrantTypes=o.join(" ")}}),g=s({get(){return(e.value.RedirectUris?e.value.RedirectUris.split(" "):[]).join(`
`)},set(o){const l=o.split(`
`);e.value.RedirectUris=l.join(" ")}}),q={style:{width:"150px"}},I={span:14},x={ClientName:[{required:!0,message:"Please input Client name",trigger:"change"},{min:3,max:25,message:"Length should be 3 to 25",trigger:"blur"}],ClientId:[{required:!0,message:"Please input Client Id",trigger:"change"},{min:3,max:25,message:"Length should be 3 to 25",trigger:"blur"}],Scopes:[{required:!0,message:"Scopes cannot be empty",trigger:"change"}],GrantTypes:[{required:!0,message:"GrantTypes cannot be empty",trigger:"change"}]};return U({validate:y,resetForm:R,model:e}),(o,l)=>{const r=u("a-form-item"),i=u("a-checkbox"),d=u("a-input"),S=u("a-image"),f=u("a-textarea"),b=u("a-select"),T=u("a-form");return N(),P(T,{ref_key:"form",ref:p,"label-col":q,model:e.value,rules:x,"wrapper-col":I},{default:n(()=>[a(r,{label:"Id"},{default:n(()=>[L("label",null,D(e.value.Id),1)]),_:1}),a(r,{label:"Enabled"},{default:n(()=>[a(i,{checked:e.value.Enabled,"onUpdate:checked":l[0]||(l[0]=t=>e.value.Enabled=t)},null,8,["checked"])]),_:1}),a(r,{label:"Client name",name:"ClientName",required:""},{default:n(()=>[a(d,{value:e.value.ClientName,"onUpdate:value":l[1]||(l[1]=t=>e.value.ClientName=t)},null,8,["value"])]),_:1}),a(r,{label:"Client Id",name:"ClientId",required:""},{default:n(()=>[a(d,{value:e.value.ClientId,"onUpdate:value":l[2]||(l[2]=t=>e.value.ClientId=t)},null,8,["value"])]),_:1}),a(r,{label:" "},{default:n(()=>[a(i,{checked:e.value.RequireConsent,"onUpdate:checked":l[3]||(l[3]=t=>e.value.RequireConsent=t)},{default:n(()=>[c(" Require consent ")]),_:1},8,["checked"]),a(i,{checked:e.value.RequireSecret,"onUpdate:checked":l[4]||(l[4]=t=>e.value.RequireSecret=t)},{default:n(()=>[c(" Require Secret ")]),_:1},8,["checked"]),a(i,{checked:e.value.RequirePkce,"onUpdate:checked":l[5]||(l[5]=t=>e.value.RequirePkce=t)},{default:n(()=>[c(" Require PKCE ")]),_:1},8,["checked"])]),_:1}),a(r,{label:"Client Url"},{default:n(()=>[a(d,{value:e.value.ClientURI,"onUpdate:value":l[6]||(l[6]=t=>e.value.ClientURI=t)},null,8,["value"])]),_:1}),a(r,{label:"Client Logo Url"},{default:n(()=>[a(S,{src:o.client.LogoURI,height:"32px",width:"32px"},null,8,["src"]),a(d,{value:e.value.LogoURI,"onUpdate:value":l[7]||(l[7]=t=>e.value.LogoURI=t)},null,8,["value"])]),_:1}),a(r,{label:"Description"},{default:n(()=>[a(f,{value:e.value.Description,"onUpdate:value":l[8]||(l[8]=t=>e.value.Description=t)},null,8,["value"])]),_:1}),a(r,{label:"Grant types",name:"GrantTypes"},{default:n(()=>[a(b,{value:m.value,"onUpdate:value":l[9]||(l[9]=t=>m.value=t),options:h.value,mode:"multiple"},null,8,["value","options"])]),_:1}),a(r,{label:"Scopes",name:"Scopes"},{default:n(()=>[a(b,{value:v.value,"onUpdate:value":l[10]||(l[10]=t=>v.value=t),options:k,mode:"multiple"},null,8,["value"])]),_:1}),a(r,{label:"Redirect Uris"},{default:n(()=>[a(f,{value:g.value,"onUpdate:value":l[11]||(l[11]=t=>g.value=t)},null,8,["value"])]),_:1})]),_:1},8,["model"])}}});export{F as _};