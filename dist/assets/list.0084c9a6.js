import{l as e,c as a,a as l,b as t,d as c,e as n}from"./cache.cd00caf8.js";import{r,z as d,G as o,_ as s,d as h,i,M as u,o as m,c as f,f as p,w as g,h as y,N as C,n as w,k as _,p as v,m as b,e as k}from"./index.1248482e.js";const N=e=>(v("data-v-7052a0aa"),e=e(),b(),e),V={class:"app-container"},x=N((()=>k("span",null,"缓存列表",-1))),K=_("刷新"),F=N((()=>k("span",null,"键名列表",-1))),R=_("刷新"),S=N((()=>k("span",null,"缓存内容",-1))),$=_("清理全部"),H=h({name:"CacheList"});var L=s(h({...H,setup(s){const{cacheNames:h,cacheKeys:_,cacheForm:v,loading:b,subLoading:k,tableHeight:N,refreshCacheNames:H,handleClearCacheName:L,getCacheKeys:O,refreshCacheKeys:U,handleClearCacheKey:j,nameFormatter:z,keyFormatter:A,handleCacheValue:G,handleClearCacheAll:I}=(()=>{const{proxy:s}=o(),h=r([]),i=r([]),u=r({}),m=r(!0),f=r(!1),p=r(""),g=r(window.innerHeight-200),y=()=>{m.value=!0,e().then((e=>{200===e.code&&(h.value=e.data,m.value=!1)}))},C=e=>{const a=e?e.cacheName:p.value;""!==a&&(f.value=!0,l(a).then((e=>{200===e.code&&(i.value=e.data,f.value=!1,p.value=a)})))};return d((()=>{y()})),{cacheNames:h,cacheKeys:i,cacheForm:u,loading:m,subLoading:f,tableHeight:g,refreshCacheNames:()=>{y(),s.$modal.msgSuccess("刷新缓存列表成功")},handleClearCacheName:e=>{a(e.cacheName).then((e=>{200===e.code&&(s.$modal.msgSuccess("清理缓存名称["+p.value+"]成功"),C())}))},getCacheKeys:C,refreshCacheKeys:()=>{C(),s.$modal.msgSuccess("刷新键名列表成功")},handleClearCacheKey:e=>{t(e).then((a=>{200===a.code&&(s.$modal.msgSuccess("清理缓存键名["+e+"]成功"),C())}))},nameFormatter:e=>e.cacheName.replace(":",""),keyFormatter:e=>e.replace(p.value,""),handleCacheValue:e=>{c(p.value,e).then((e=>{200===e.code&&(u.value=e.data)}))},handleClearCacheAll:()=>{n().then((e=>{200===e.code&&s.$modal.msgSuccess("清理全部缓存成功")}))}}})();return(e,a)=>{const l=i("el-link"),t=i("el-table-column"),c=i("el-table"),n=i("el-card"),r=i("el-col"),d=i("el-input"),o=i("el-form-item"),s=i("el-row"),M=i("el-form"),q=u("loading");return m(),f("div",V,[p(s,{gutter:10},{default:g((()=>[p(r,{span:8},{default:g((()=>[p(n,{style:{height:"calc(100vh - 125px)"}},{header:g((()=>[x,p(l,{class:"cache_refresh",underline:!1,type:"primary",icon:"Refresh",onClick:a[0]||(a[0]=e=>y(H)()),title:"点击刷新缓存"},{default:g((()=>[K])),_:1})])),default:g((()=>[C((m(),w(c,{data:y(h),height:y(N),"highlight-current-row":"",onRowClick:y(O),style:{width:"100%"}},{default:g((()=>[p(t,{align:"center",label:"序号",width:"60",type:"index"}),p(t,{label:"缓存名称",align:"center",prop:"cacheName","show-overflow-tooltip":!0,formatter:y(z)},null,8,["formatter"]),p(t,{label:"备注",align:"center",prop:"remark","show-overflow-tooltip":!0}),p(t,{label:"操作",width:"90",align:"center","class-name":"small-padding fixed-width"},{default:g((e=>[p(l,{type:"danger",icon:"delete",underline:!1,onClick:a=>y(L)(e.row),title:"点击删除缓存"},null,8,["onClick"])])),_:1})])),_:1},8,["data","height","onRowClick"])),[[q,y(b)]])])),_:1})])),_:1}),p(r,{span:8},{default:g((()=>[p(n,{style:{height:"calc(100vh - 125px)"}},{header:g((()=>[F,p(l,{class:"cache_refresh",type:"primary",icon:"refresh",underline:!1,onClick:a[1]||(a[1]=e=>y(U)()),title:"点击刷新缓存"},{default:g((()=>[R])),_:1})])),default:g((()=>[C((m(),w(c,{data:y(_),height:y(N),"highlight-current-row":"",onRowClick:y(G),style:{width:"100%"}},{default:g((()=>[p(t,{align:"center",label:"序号",width:"60",type:"index"}),p(t,{label:"缓存键名",align:"center","show-overflow-tooltip":!0,formatter:y(A)},null,8,["formatter"]),p(t,{label:"操作",width:"90",align:"center","class-name":"small-padding fixed-width"},{default:g((e=>[p(l,{type:"warning",icon:"delete",underline:!1,onClick:a=>y(j)(e.row),title:"点击删除缓存"},null,8,["onClick"])])),_:1})])),_:1},8,["data","height","onRowClick"])),[[q,y(k)]])])),_:1})])),_:1}),p(r,{span:8},{default:g((()=>[p(n,{bordered:!1,style:{height:"calc(100vh - 125px)"}},{header:g((()=>[S,p(l,{class:"cache_refresh",type:"primary",icon:"delete",underline:!1,onClick:a[2]||(a[2]=e=>y(I)())},{default:g((()=>[$])),_:1})])),default:g((()=>[p(M,{model:y(v)},{default:g((()=>[p(s,{gutter:32},{default:g((()=>[p(r,{offset:1,span:22},{default:g((()=>[p(o,{label:"缓存名称:",prop:"cacheName"},{default:g((()=>[p(d,{modelValue:y(v).cacheName,"onUpdate:modelValue":a[3]||(a[3]=e=>y(v).cacheName=e),readOnly:!0},null,8,["modelValue"])])),_:1})])),_:1}),p(r,{offset:1,span:22},{default:g((()=>[p(o,{label:"缓存键名:",prop:"cacheKey"},{default:g((()=>[p(d,{modelValue:y(v).cacheKey,"onUpdate:modelValue":a[4]||(a[4]=e=>y(v).cacheKey=e),readOnly:!0},null,8,["modelValue"])])),_:1})])),_:1}),p(r,{offset:1,span:22},{default:g((()=>[p(o,{label:"缓存内容:",prop:"cacheValue"},{default:g((()=>[p(d,{modelValue:y(v).cacheValue,"onUpdate:modelValue":a[5]||(a[5]=e=>y(v).cacheValue=e),type:"textarea",readOnly:!0,autosize:{minRows:15}},null,8,["modelValue"])])),_:1})])),_:1})])),_:1})])),_:1},8,["model"])])),_:1})])),_:1})])),_:1})])}}}),[["__scopeId","data-v-7052a0aa"]]);export{L as default};