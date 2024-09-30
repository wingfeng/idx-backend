import{O as U}from"./index-B0yw5fqV.mjs";import{a6 as f,aG as s,aH as y,aI as h,aQ as m,aR as V,aa as n,aS as A,aT as P,aU as R,aV as S,aJ as l,ao as b,aW as a,F as _,aX as g,aK as t,h as d,m as v,aY as x}from"../jse/index-index-CqEm-tRS.mjs";import{P as z,V as F,a as C}from"./bootstrap-t35vpcz_.mjs";const p=f({__name:"link",props:{class:{default:""},href:{default:""},asChild:{type:Boolean},as:{default:"a"}},setup(c){const o=c;return(e,u)=>(s(),y(n(U),{as:e.as,"as-child":e.asChild,class:V(n(A)("text-primary hover:text-primary-hover",o.class)),href:e.href,target:"_blank"},{default:h(()=>[m(e.$slots,"default")]),_:3},8,["as","as-child","class","href"]))}}),k=f({name:"RenderContent",__name:"render-content",props:{content:{},props:{default:()=>({})}},setup(c){const o=c,e=typeof o.content=="object"&&o.content!==null;return(u,j)=>n(e)?(s(),y(S(u.content),P(R({key:0},o)),null,16)):n(e)?g("",!0):(s(),l(_,{key:1},[b(a(u.content),1)],64))}}),L={class:"relative h-full"},O={key:0,class:"bg-card px-6 py-4"},q={key:0,class:"mb-2 flex justify-between text-lg font-semibold"},G={key:0,class:"text-muted-foreground"},M={key:1,class:"bg-card align-center absolute bottom-0 left-0 right-0 flex px-6 py-4"},H=f({name:"Page",__name:"page",props:{title:{default:""},description:{default:""},contentClass:{default:""},showFooter:{type:Boolean,default:!1}},setup(c){const o=c;return(e,u)=>(s(),l("div",L,[e.description||e.$slots.description||e.title?(s(),l("div",O,[m(e.$slots,"title",{},()=>[e.title?(s(),l("div",q,[b(a(e.title)+" ",1),m(e.$slots,"extra")])):g("",!0)]),m(e.$slots,"description",{},()=>[e.description?(s(),l("p",G,a(e.description),1)):g("",!0)])])):g("",!0),t("div",{class:V([e.contentClass,"m-4"])},[m(e.$slots,"default")],2),o.showFooter?(s(),l("div",M,[m(e.$slots,"footer")])):g("",!0)]))}});var W={authorEmail:"wingfeng@gmail.com",authorName:"wingfeng",authorUrl:"https://github.com/wingfeng",buildTime:"2024-09-30 23:06:01",dependencies:{jsonwebtoken:"^9.0.2",nitropack:"^2.9.7","@vben/access":"5.2.1","@vben/common-ui":"5.2.1","@vben/constants":"5.2.1","@vben/hooks":"5.2.1","@vben/icons":"5.2.1","@vben/layouts":"5.2.1","@vben/locales":"5.2.1","@vben/plugins":"5.2.1","@vben/preferences":"5.2.1","@vben/request":"5.2.1","@vben/stores":"5.2.1","@vben/styles":"5.2.1","@vben/types":"5.2.1","@vben/utils":"5.2.1","@vueuse/core":"^11.0.3","ant-design-vue":"^4.2.3",dayjs:"^1.11.13",pinia:"2.2.2",vue:"^3.4.38","vue-router":"^4.4.3","@commitlint/cli":"^19.4.1","@commitlint/config-conventional":"^19.4.1","@vben/node-utils":"5.2.1","commitlint-plugin-function-rules":"^4.0.0","cz-git":"^1.9.4",czg:"^1.9.4","eslint-config-turbo":"^2.1.0","eslint-plugin-command":"^0.2.3","eslint-plugin-import-x":"^4.1.1",prettier:"^3.3.3","prettier-plugin-tailwindcss":"^0.6.6","@stylistic/stylelint-plugin":"^3.0.1","stylelint-config-recess-order":"^5.1.0","stylelint-scss":"^6.5.1","@changesets/git":"^3.0.0","@manypkg/get-packages":"^2.2.2",chalk:"^5.3.0",consola:"^3.2.3",execa:"^9.3.1","find-up":"^7.0.0",nanoid:"^5.0.7",ora:"^8.1.0","pkg-types":"^1.2.0",rimraf:"^6.0.1","@iconify/json":"^2.2.242","@iconify/tailwind":"^1.1.3","@tailwindcss/nesting":"0.0.0-insiders.565cd3e","@tailwindcss/typography":"^0.5.15",autoprefixer:"^10.4.20",cssnano:"^7.0.5",postcss:"^8.4.41","postcss-antd-fixes":"^0.2.0","postcss-import":"^16.1.0","postcss-preset-env":"^10.0.2",tailwindcss:"^3.4.10","tailwindcss-animate":"^1.0.7",vite:"^5.4.2","@intlify/unplugin-vue-i18n":"^4.0.0","@jspm/generator":"^2.1.3",archiver:"^7.0.1",cheerio:"1.0.0","get-port":"^7.1.0","html-minifier-terser":"^7.2.0","resolve.exports":"^2.0.2","vite-plugin-lib-inject-css":"^2.1.1","vite-plugin-pwa":"^0.20.2","vite-plugin-vue-devtools":"^7.3.9","@iconify/vue":"^4.1.2","lucide-vue-next":"^0.436.0","@ctrl/tinycolor":"^4.1.0","@tanstack/vue-store":"^0.5.5","@vue/shared":"^3.4.38",clsx:"^2.1.1",defu:"^6.1.4","lodash.clonedeep":"^4.5.0",nprogress:"^0.2.0","tailwind-merge":"^2.5.2","theme-colors":"^0.1.0","@vben-core/shared":"5.2.1","radix-vue":"^1.9.5",sortablejs:"^1.15.2","@vben-core/typings":"5.2.1","@vben-core/composables":"5.2.1","@vben-core/icons":"5.2.1","@vben-core/shadcn-ui":"5.2.1","@radix-icons/vue":"^1.0.0","class-variance-authority":"^0.7.0","@vben-core/popup-ui":"5.1.1","@vueuse/integrations":"^11.0.3",qrcode:"^1.5.4","watermark-js-plus":"^1.5.4","@vben-core/layout-ui":"5.2.1","@vben-core/menu-ui":"5.2.1","@vben-core/tabs-ui":"5.2.1",echarts:"^5.5.1",axios:"^1.7.5","@intlify/core-base":"^9.14.0","vue-i18n":"^9.14.0","@vben-core/preferences":"5.2.1","pinia-plugin-persistedstate":"^3.2.3","@vben-core/design":"5.2.1","@clack/prompts":"^0.7.0",cac:"^6.7.14","circular-dependency-scanner":"^2.2.2",depcheck:"^1.4.7",publint:"^0.2.10"},devDependencies:{"@types/jsonwebtoken":"^9.0.6",h3:"^1.12.0","@eslint/js":"^9.9.1","@types/eslint":"^9.6.1","@typescript-eslint/eslint-plugin":"^8.3.0","@typescript-eslint/parser":"^8.3.0",eslint:"^9.9.1","eslint-config-prettier":"^9.1.0","eslint-plugin-eslint-comments":"^3.2.0","eslint-plugin-jsdoc":"^50.2.2","eslint-plugin-jsonc":"^2.16.0","eslint-plugin-n":"^17.10.2","eslint-plugin-no-only-tests":"^3.3.0","eslint-plugin-perfectionist":"^3.3.0","eslint-plugin-prettier":"^5.2.1","eslint-plugin-regexp":"^2.6.0","eslint-plugin-unicorn":"^55.0.0","eslint-plugin-unused-imports":"^4.1.3","eslint-plugin-vitest":"^0.5.4","eslint-plugin-vue":"^9.27.0",globals:"^15.9.0","jsonc-eslint-parser":"^2.4.0","vue-eslint-parser":"^9.4.3",postcss:"^8.4.41","postcss-html":"^1.7.0","postcss-scss":"^4.0.9",prettier:"^3.3.3",stylelint:"^16.9.0","stylelint-config-recommended":"^14.0.1","stylelint-config-recommended-scss":"^14.1.0","stylelint-config-recommended-vue":"^1.5.0","stylelint-config-standard":"^36.0.1","stylelint-order":"^6.0.4","stylelint-prettier":"^5.0.2","@types/chalk":"^2.2.0","@types/postcss-import":"^14.0.3","@vben/node-utils":"5.2.1","@types/archiver":"^6.0.2","@types/html-minifier-terser":"^7.0.2","@vitejs/plugin-vue":"^5.1.3","@vitejs/plugin-vue-jsx":"^4.0.1",dayjs:"^1.11.13",dotenv:"^16.4.5",rollup:"^4.21.2","rollup-plugin-visualizer":"^5.12.0",sass:"^1.77.8",vite:"^5.4.2","vite-plugin-compression":"^0.5.1","vite-plugin-dts":"4.0.3","vite-plugin-html":"^3.2.2","@types/lodash.clonedeep":"^4.5.9","@types/nprogress":"^0.2.3","@types/sortablejs":"^1.15.8","@types/qrcode":"^1.5.5","axios-mock-adapter":"^2.0.0"},homepage:"",license:"MIT",version:"0.1.1"};const J={class:"text-foreground mt-3 text-sm leading-6"},K={class:"card-box p-5"},Q=t("div",null,[t("h5",{class:"text-foreground text-lg"},"基本信息")],-1),X={class:"mt-4"},Y={class:"grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4"},Z={class:"text-foreground text-sm font-medium leading-6"},ee={class:"text-foreground mt-1 text-sm leading-6 sm:mt-2"},te={class:"card-box mt-6 p-5"},se=t("div",null,[t("h5",{class:"text-foreground text-lg"},"生产环境依赖")],-1),ne={class:"mt-4"},oe={class:"grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4"},ie={class:"text-foreground text-sm"},le={class:"text-foreground/80 mt-1 text-sm sm:mt-2"},re={class:"card-box mt-6 p-5"},ae=t("div",null,[t("h5",{class:"text-foreground text-lg"},"开发环境依赖")],-1),ce={class:"mt-4"},de={class:"grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4"},pe={class:"text-foreground text-sm"},ue={class:"text-foreground/80 mt-1 text-sm sm:mt-2"},me=f({name:"AboutUI",__name:"about",props:{description:{default:"是一个现代化开箱即用的中后台解决方案，采用最新的技术栈，包括 Vue 3.0、Vite、TailwindCSS 和 TypeScript 等前沿技术，代码规范严谨，提供丰富的配置选项，旨在为中大型项目的开发提供现成的开箱即用解决方案及丰富的示例，同时，它也是学习和深入前端技术的一个极佳示例。"},name:{default:"Vben Admin"},title:{default:"关于项目"}},setup(c){const{authorEmail:o,authorName:e,authorUrl:u,buildTime:j,dependencies:w={},devDependencies:$={},homepage:T,license:B,version:D}=W||{},I=[{content:D,title:"版本号"},{content:B,title:"开源许可协议"},{content:j,title:"最后构建时间"},{content:d(p,{href:T},{default:()=>"点击查看"}),title:"主页"},{content:d(p,{href:z},{default:()=>"点击查看"}),title:"文档地址"},{content:d(p,{href:F},{default:()=>"点击查看"}),title:"预览地址"},{content:d(p,{href:C},{default:()=>"点击查看"}),title:"Github"},{content:d("div",[d(p,{class:"mr-2",href:u},{default:()=>e}),d(p,{href:`mailto:${o}`},{default:()=>o})]),title:"作者"}],N=Object.keys(w).map(r=>({content:w[r],title:r})),E=Object.keys($).map(r=>({content:$[r],title:r}));return(r,ge)=>(s(),y(n(H),{title:r.title},{description:h(()=>[t("p",J,[v(n(p),{href:n(C)},{default:h(()=>[b(a(r.name),1)]),_:1},8,["href"]),b(" "+a(r.description),1)])]),default:h(()=>[t("div",K,[Q,t("div",X,[t("dl",Y,[(s(),l(_,null,x(I,i=>t("div",{key:i.title,class:"border-border border-t px-4 py-6 sm:col-span-1 sm:px-0"},[t("dt",Z,a(i.title),1),t("dd",ee,[v(n(k),{content:i.content},null,8,["content"])])])),64))])])]),t("div",te,[se,t("div",ne,[t("dl",oe,[(s(!0),l(_,null,x(n(N),i=>(s(),l("div",{key:i.title,class:"border-border border-t px-4 py-3 sm:col-span-1 sm:px-0"},[t("dt",ie,a(i.title),1),t("dd",le,[v(n(k),{content:i.content},null,8,["content"])])]))),128))])])]),t("div",re,[ae,t("div",ce,[t("dl",de,[(s(!0),l(_,null,x(n(E),i=>(s(),l("div",{key:i.title,class:"border-border border-t px-4 py-3 sm:col-span-1 sm:px-0"},[t("dt",pe,a(i.title),1),t("dd",ue,[v(n(k),{content:i.content},null,8,["content"])])]))),128))])])])]),_:1},8,["title"]))}}),_e=f({name:"About",__name:"index",setup(c){return(o,e)=>(s(),y(n(me)))}});export{_e as default};