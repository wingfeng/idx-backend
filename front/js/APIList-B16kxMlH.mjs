var d=(l,E,s)=>new Promise((m,t)=>{var D=o=>{try{v(s.next(o))}catch(g){t(g)}},I=o=>{try{v(s.throw(o))}catch(g){t(g)}},v=o=>o.done?m(o.value):Promise.resolve(o.value).then(D,I);v((s=s.apply(l,E)).next())});import{u as le,R as ne,M as oe}from"./index.es-BVwDNJs5.mjs";import{r as U,S as ie,D as w,E as re,e as se,M as ue}from"./bootstrap-t35vpcz_.mjs";import{_ as ce}from"./APIForm.vue_vue_type_script_setup_true_lang-CeTtoO_a.mjs";import{a6 as de,P as p,J as z,a7 as pe,aF as r,aG as M,aJ as me,m as a,aI as n,aa as i,ao as N,h as A,aH as R,aX as ve}from"../jse/index-index-CqEm-tRS.mjs";const $={Page:"/api/v1/oauth2/api/page",Get:"/api/v1/oauth2/api/get",PUT:"/api/v1/oauth2/api",Delete:"/api/v1/oauth2/api/del"};function fe(l){return d(this,null,function*(){return U.post($.Page,l)})}function _e(l){return d(this,null,function*(){return U.delete($.Delete,{params:{id:l}})})}function ye(l){return d(this,null,function*(){return U.put($.PUT,l)})}function ge(l){return d(this,null,function*(){return fe(l)})}function he(l){return d(this,null,function*(){return ye(l)})}function ke(l){return d(this,null,function*(){return _e(l)})}const Ne={class:"p-5"},Se=de({__name:"APIList",setup(l){const E=[{title:"Id",dataIndex:"Id",key:"Id",sorter:!0,defaultSortOrder:"ascend"},{title:"Enabled",dataIndex:"Enabled",key:"Enabled",sorter:!0},{title:"Name",dataIndex:"Name",key:"Name",name:"Name",sorter:!0,filtered:!0},{title:"DisplayName",dataIndex:"DisplayName",key:"DisplayName",sorter:!0,filtered:!0},{title:"Description",dataIndex:"Description",key:"Description",filtered:!0},{title:"Actions",key:"action"}],s=p("Id"),m=p("asc"),t=p({Name:"",DisplayName:""}),D=z(()=>{const e=[];return t.value.Name!==""&&e.push("Name like ?"),t.value.DisplayName!==""&&e.push("Display_Name like ?"),e}),I=z(()=>{const e=[];return t.value.Name!==""&&e.push(`%${t.value.Name}%`),t.value.DisplayName!==""&&e.push(`%${t.value.DisplayName}%`),e}),{data:v,run:o,loading:g,total:B,current:b,pageSize:C}=le(ge,{defaultParams:{page:1,pageSize:10,sortField:"Id",sortOrder:"asc",filters:D,args:I},pagination:{currentKey:"page",pageSizeKey:"pageSize"}}),G={Id:"Id",DisplayName:"display_name"},L=z(()=>({total:B.value,current:b.value,pageSize:C.value})),T=p(),h=()=>{setTimeout(()=>{o({page:b.value,pageSize:C.value,sortField:s.value,sortOrder:m.value,filters:D,args:I})},500)},V=()=>{console.log("searchModel",t.value),h()},J=()=>{t.value={Name:"",DisplayName:""},T.value.resetFields(),h()},f=p(!1),P=p(),O=p(),K=e=>{P.value=e,f.value=!0},q=e=>{ue.confirm({title:`Deleting API ${e}`,content:`Are you sure delete this API ${e}?`,onOk(){ke(e),h()},onCancel(){console.log("Cancel")}}),console.log("record",e)},H=()=>d(this,null,function*(){try{const e=yield O.value.validate();if(!e)return;console.log("form validate",e)}catch(e){console.log("error submit:",e);return}he(P.value),f.value=!1,h()}),X=()=>{O.value.resetForm(),f.value=!1},j=()=>{P.value={Id:0,Enabled:!0,Name:"",DisplayName:"",Description:""},f.value=!0},Q=(e,u,_)=>{if(b.value=e.current,C.value=e.pageSize,_.field){const y=G[_.field];s.value=y!=null?y:_.field,m.value=_.order}else s.value="Id",m.value="asc";console.log("filter",u),o({page:b.value,pageSize:C.value,sortField:s.value,sortOrder:m.value,filter:u})};return pe(()=>{h()}),(e,u)=>{const _=r("a-page-header"),y=r("a-input"),x=r("a-form-item"),k=r("a-button"),W=r("a-form"),Y=r("a-divider"),Z=r("a-space"),ee=r("a-checkbox"),ae=r("a-table"),te=r("a-modal");return M(),me("div",Ne,[a(_,{style:{border:"1px solid rgb(235 237 240)"},"sub-title":"API list page",title:"APIs"}),a(W,{ref_key:"searchForm",ref:T,model:t.value,layout:"inline"},{default:n(()=>[a(x,{label:"Name"},{default:n(()=>[a(y,{value:t.value.Name,"onUpdate:value":u[0]||(u[0]=c=>t.value.Name=c),placeholder:"Name"},null,8,["value"])]),_:1}),a(x,{label:"DisplayName"},{default:n(()=>[a(y,{value:t.value.DisplayName,"onUpdate:value":u[1]||(u[1]=c=>t.value.DisplayName=c),placeholder:"DisplayName"},null,8,["value"])]),_:1}),a(x,null,{default:n(()=>[a(k,{type:"primary",onClick:V},{icon:n(()=>[a(i(ie))]),default:n(()=>[N(" Search ")]),_:1}),a(k,{icon:A(i(ne)),type:"primary",onClick:J},{default:n(()=>[N(" Reset ")]),_:1},8,["icon"]),a(i(w),{type:"vertical"}),a(k,{icon:A(i(oe)),type:"primary",onClick:j},{default:n(()=>[N(" New ")]),_:1},8,["icon"])]),_:1})]),_:1},8,["model"]),a(Y,{type:"horizontal"}),a(ae,{columns:E,"data-source":i(v)?i(v).list:null,loading:i(g),pagination:L.value,run:i(o),onChange:Q},{bodyCell:n(({column:c,record:S})=>[c.key==="action"?(M(),R(Z,{key:0},{default:n(()=>[a(k,{icon:A(i(re)),type:"primary",onClick:F=>K(S)},{default:n(()=>[N(" Edit ")]),_:2},1032,["icon","onClick"]),a(i(w),{type:"vertical"}),a(k,{icon:A(i(se)),danger:"",onClick:F=>q(S.Id)},{default:n(()=>[N(" Delete ")]),_:2},1032,["icon","onClick"])]),_:2},1024)):c.key==="Enabled"?(M(),R(ee,{key:1,checked:S.Enabled,"onUpdate:checked":F=>S.Enabled=F},null,8,["checked","onUpdate:checked"])):ve("",!0)]),_:1},8,["data-source","loading","pagination","run"]),a(te,{open:f.value,"onUpdate:open":u[2]||(u[2]=c=>f.value=c),title:"Edit Client Info",onCancel:X,onOk:H},{default:n(()=>[a(ce,{ref_key:"modalForm",ref:O,"form-model":P.value},null,8,["form-model"])]),_:1},8,["open"])])}}});export{Se as default};
