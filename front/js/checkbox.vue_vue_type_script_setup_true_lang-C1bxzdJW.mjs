var u=Object.getOwnPropertySymbols;var g=Object.prototype.hasOwnProperty,_=Object.prototype.propertyIsEnumerable;var p=(a,o)=>{var s={};for(var e in a)g.call(a,e)&&o.indexOf(e)<0&&(s[e]=a[e]);if(a!=null&&u)for(var e of u(a))o.indexOf(e)<0&&_.call(a,e)&&(s[e]=a[e]);return s};import{x as h,M as v,k as x}from"./index-6Td_2BJY.mjs";import{a6 as k,J as w,aG as y,aH as C,aI as m,m as i,aa as t,aQ as B,aU as b,aS as M,aZ as f,a_ as S,aJ as $,aK as q,F}from"../jse/index-index-CWuPeoDv.mjs";import{r as J}from"./input-password.vue_vue_type_script_setup_true_lang-DUiXJGVf.mjs";const N=k({__name:"Checkbox",props:{class:{},defaultChecked:{type:Boolean},checked:{type:[Boolean,String]},disabled:{type:Boolean},required:{type:Boolean},name:{},value:{},id:{},asChild:{type:Boolean},as:{}},emits:["update:checked"],setup(a,{emit:o}){const s=a,e=o,c=w(()=>{const l=s,{class:r}=l;return p(l,["class"])}),n=h(c,e);return(r,d)=>(y(),C(t(x),b(t(n),{class:t(M)("focus-visible:ring-ring data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground border-border peer h-4 w-4 shrink-0 rounded-sm border focus-visible:outline-none focus-visible:ring-1 disabled:cursor-not-allowed disabled:opacity-50",s.class)}),{default:m(()=>[i(t(v),{class:"flex h-full w-full items-center justify-center text-current"},{default:m(()=>[B(r.$slots,"default",{},()=>[i(t(J),{class:"h-4 w-4"})])]),_:3})]),_:3},16,["class"]))}}),P=["for"],G=k({__name:"checkbox",props:f({name:{},defaultChecked:{type:Boolean},checked:{type:[Boolean,String]},disabled:{type:Boolean},required:{type:Boolean},value:{},id:{},asChild:{type:Boolean},as:{}},{checked:{type:Boolean},checkedModifiers:{}}),emits:f(["update:checked"],["update:checked"]),setup(a,{emit:o}){const s=a,e=o,c=S(a,"checked"),n=h(s,e);return(r,d)=>(y(),$(F,null,[i(t(N),b(t(n),{id:r.name,checked:c.value,"onUpdate:checked":d[0]||(d[0]=l=>c.value=l)}),null,16,["id","checked"]),q("label",{for:r.name,class:"ml-2 cursor-pointer text-sm"},[B(r.$slots,"default")],8,P)],64))}});export{G as _};
