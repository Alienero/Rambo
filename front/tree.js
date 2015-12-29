var nodes = []
var links = []
var depths_nodes = []

function addNode(node,parent){

}

var option = {
  title : {
    text: 'Force',
    subtext: 'Force-directed tree',
    x:'right',
    y:'bottom'
  },
  tooltip : {
    trigger: 'item',
    formatter: '{a} : {b}'
  },
  toolbox: {
    show : true,
    feature : {
      restore : {show: true},
      saveAsImage : {show: true}
    }
  },
  legend: {
    x: 'left',
    data:['叶子节点','非叶子节点', '根节点']
  },
  series : [
    {
      type:'force',
      name : "SQL AST Tree",
      categories : [
        {
          name: '叶子节点',
          itemStyle: {
            normal: {
              color : '#ff7f50'
            }
          }
        },
        {
          name: '非叶子节点',
          itemStyle: {
            normal: {
              color : '#6f57bc'
            }
          }
        },
        {
          name: '根节点',
          itemStyle: {
            normal: {
              color : '#af0000'
            }
          }
        }
      ],
      itemStyle: {
        normal: {
          label: {
            show: false
          },
          nodeStyle : {
            brushType : 'both',
            strokeColor : 'rgba(255,215,0,0.6)',
            lineWidth : 1
          }
        }
      },
      // minRadius : 2,
      // maxRadius : 10,
      nodes : nodes,
      links : links
    }
  ]
};
