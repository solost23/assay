import{I as a}from"./index.1248482e.js";const e=async()=>await a({url:"/monitor/cache",method:"get"}),t=async()=>await a({url:"/monitor/cache/getNames",method:"get"}),c=async e=>await a({url:"/monitor/cache/getKeys/"+e,method:"get"}),o=async(e,t)=>await a({url:"/monitor/cache/getValue/"+e+"/"+t,method:"get"}),r=async e=>await a({url:"/monitor/cache/clearCacheName/"+e,method:"delete"}),s=async e=>await a({url:"/monitor/cache/clearCacheKey/"+e,method:"delete"}),m=async()=>await a({url:"/monitor/cache/clearCacheAll",method:"delete"});export{c as a,s as b,r as c,o as d,m as e,e as g,t as l};
