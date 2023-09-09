<script>

  let credit = 10000
  let min_bet = 0
  let totalbet = 0
  let bet_0 = 0
  let bet_1 = 0
  let bet_2 = 0
  let bet_3 = 0
  let sound = 0;
  let flag_minimalbet = false;
  let path_card = "CARD/WHITE/";
  let spin = [
    new Audio("/sounds/spin.mp3"),
  ];
  let win = [
    new Audio("/sounds/win.mp3"),
    new Audio("/sounds/win.mp3"),
  ];
  let list_point = [
    {id:"1",code:"RF",name:"Royal Flush",poin:50000},
    {id:"2",code:"5K",name:"5 Of A Kind",poin:20000},
    {id:"3",code:"SF",name:"Straight Flush",poin:12000},
    {id:"4",code:"4K",name:"4 Of A Kind",poin:1200},
    {id:"5",code:"FH",name:"Full House",poin:700},
    {id:"6",code:"FL",name:"Flush",poin:500},
    {id:"7",code:"ST",name:"Straight",poin:300},
    {id:"8",code:"3K",name:"3 Of A Kind",poin:200},
    {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:100},
    {id:"10",code:"AP",name:"Ace Pair",poin:100},
  ]
  const card_result_data = [
    {id:"2_diamond",val:"2",val_display:2,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_2.png"},
    {id:"3_diamond",val:"3",val_display:3,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_3.png"},
    {id:"4_diamond",val:"4",val_display:4,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_4.png"},
    {id:"5_diamond",val:"5",val_display:5,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_5.png"},
    {id:"6_diamond",val:"6",val_display:6,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_6.png"},
    {id:"7_diamond",val:"7",val_display:7,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_7.png"},
    {id:"8_diamond",val:"8",val_display:8,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_8.png"},
    {id:"9_diamond",val:"9",val_display:9,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_9.png"},
    {id:"10_diamond",val:"10",val_display:10,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_10.png"},
    {id:"j_diamond",val:"J",val_display:11,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_J.png"},
    {id:"q_diamond",val:"Q",val_display:12,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_Q.png"},
    {id:"k_diamond",val:"K",val_display:13,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_K.png"},
    {id:"as_diamond",val:"AS",val_display:14,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_AS.png"},
    {id:"2_heart",val:"2",val_display:2,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_2.png"},
    {id:"3_heart",val:"3",val_display:3,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_3.png"},
    {id:"4_heart",val:"4",val_display:4,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_4.png"},
    {id:"5_heart",val:"5",val_display:5,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_5.png"},
    {id:"6_heart",val:"6",val_display:6,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_6.png"},
    {id:"7_heart",val:"7",val_display:7,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_7.png"},
    {id:"8_heart",val:"8",val_display:8,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_8.png"},
    {id:"9_heart",val:"9",val_display:9,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_9.png"},
    {id:"10_heart",val:"10",val_display:10,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_10.png"},
    {id:"j_heart",val:"J",val_display:11,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_J.png"},
    {id:"q_heart",val:"Q",val_display:12,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_Q.png"},
    {id:"k_heart",val:"K",val_display:13,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_K.png"},
    {id:"as_heart",val:"AS",val_display:14,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_AS.png"},
    {id:"2_club",val:"2",val_display:2,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_2.png"},
    {id:"3_club",val:"3",val_display:3,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_3.png"},
    {id:"4_club",val:"4",val_display:4,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_4.png"},
    {id:"5_club",val:"5",val_display:5,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_5.png"},
    {id:"6_club",val:"6",val_display:6,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_6.png"},
    {id:"7_club",val:"7",val_display:7,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_7.png"},
    {id:"8_club",val:"8",val_display:8,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_8.png"},
    {id:"9_club",val:"9",val_display:9,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_9.png"},
    {id:"10_club",val:"10",val_display:10,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_10.png"},
    {id:"j_club",val:"J",val_display:11,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_J.png"},
    {id:"q_club",val:"Q",val_display:12,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_Q.png"},
    {id:"k_club",val:"K",val_display:13,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_K.png"},
    {id:"as_club",val:"AS",val_display:14,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_AS.png"},
    {id:"2_spade",val:"2",val_display:2,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_2.png"},
    {id:"3_spade",val:"3",val_display:3,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_3.png"},
    {id:"4_spade",val:"4",val_display:4,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_4.png"},
    {id:"5_spade",val:"5",val_display:5,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_5.png"},
    {id:"6_spade",val:"6",val_display:6,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_6.png"},
    {id:"7_spade",val:"7",val_display:7,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_7.png"},
    {id:"8_spade",val:"8",val_display:8,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_8.png"},
    {id:"9_spade",val:"9",val_display:9,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_9.png"},
    {id:"10_spade",val:"10",val_display:10,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_10.png"},
    {id:"j_spade",val:"J",val_display:11,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_J.png"},
    {id:"q_spade",val:"Q",val_display:12,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_Q.png"},
    {id:"k_spade",val:"K",val_display:13,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_K.png"},
    {id:"as_spade",val:"AS",val_display:14,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_AS.png"},
    {id:"jk_black",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_BLACK.png"},
    {id:"jk_red",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_RED.png"},
  ]
  
  const pattern_stright_1 = [2,3,4,5,6]
  const pattern_stright_2 = [3,4,5,6,7]
  const pattern_stright_3 = [4,5,6,7,8]
  const pattern_stright_4 = [5,6,7,8,9]
  const pattern_stright_5 = [6,7,8,9,10]
  const pattern_stright_6 = [7,8,9,10,11]
  const pattern_stright_7 = [8,9,10,11,12]
  const pattern_stright_8 = [9,10,11,12,13]
  const pattern_stright_9 = [10,11,12,13,14]
  const pattern_stright_10 = [14,2,3,4,5]
  
  let flag_all = true
  let flag_bet = true
  let flag_fullbet = true
  let flag_deal = false
  let count_bet = 0;
  let point_result = "";
  let point_style_result = "";
  let list_point_id = "";
  let list_point_style = "";
  let card_background = "bg-transparent"
  let card_result_0_id = "NULL"
  let card_result_1_id = "NULL"
  let card_result_2_id = "NULL"
  let card_result_3_id = "NULL"
  let card_result_4_id = "NULL"
  let card_result_5_id = "NULL"
  let card_result_6_id = "NULL"
  let card_result_0_img = "./"+path_card+"CARD_FLOP.png"
  let card_result_1_img = "./"+path_card+"CARD_FLOP.png"
  let card_result_2_img = "./"+path_card+"CARD_FLOP.png"
  let card_result_3_img = "./"+path_card+"CARD_FLOP.png"
  let card_result_4_img = "./"+path_card+"CARD_FLOP.png"
  let card_result_5_img = "./"+path_card+"CARD_FLOP.png"
  let card_result_6_img = "./"+path_card+"CARD_FLOP.png"
  let card_result_0_val = ""
  let card_result_1_val = ""
  let card_result_2_val = ""
  let card_result_3_val = ""
  let card_result_4_val = ""
  let card_result_5_val = ""
  let card_result_6_val = ""
  let card_result_array_id = []
  let card_result_array_val = []
  let info_result = "";
  let info_card = [];
  let flag_win = false
  let shuffleArray = [];
  let usedIndexes = [];
  const call_play = () => {
    flag_minimalbet = false;
    bet_0 = 0
    bet_1 = 0
    bet_2 = 0
    bet_3 = 0
    count_bet = 0;
    totalbet = 0;
    flag_bet = true;
    flag_fullbet = true;
    flag_deal = false;
    flag_win = false;
    shuffleArray = [];
    usedIndexes = [];
    card_result_array_id = []
    card_result_array_val = []
    info_result = "";
    info_card = [];
    list_point_style = ""
    list_point_id = ""
    point_result = "";
    point_style_result = "";
    card_result_0_img = "./"+path_card+"CARD_FLOP.png"
    card_result_1_img = "./"+path_card+"CARD_FLOP.png"
    card_result_2_img = "./"+path_card+"CARD_FLOP.png"
    card_result_3_img = "./"+path_card+"CARD_FLOP.png"
    card_result_4_img = "./"+path_card+"CARD_FLOP.png"
    card_result_5_img = "./"+path_card+"CARD_FLOP.png"
    card_result_6_img = "./"+path_card+"CARD_FLOP.png"
    card_result_0_val = ""
    card_result_1_val = ""
    card_result_2_val = ""
    card_result_3_val = ""
    card_result_4_val = ""
    card_result_5_val = ""
    card_result_6_val = ""
	};
  const call_bet = () => {
    flag_minimalbet = true;
    if(parseInt(min_bet) > 0){
      if(credit > 0){
        if(parseInt(credit) > parseInt(min_bet)){
          count_bet = count_bet + 1
          totalbet = totalbet + 1
          sound = 0;
          spin[sound].play();
          switch(count_bet){
            case 1:
              shuffleArray_card(card_result_data)
              flag_deal = true;
              flag_fullbet = false;
              bet_0 = min_bet;
              credit = credit - bet_0;
              point_style_result = "text-red-600 font-bold"
              point_result = "-" + (totalbet * min_bet)
              break;
            case 2:
              bet_1 = min_bet;
              credit = credit - bet_1;
              point_style_result = "text-red-600 font-bold"
              point_result = "-" + (totalbet * min_bet)
              break;
            case 3:
              bet_2 = min_bet;
              credit = credit - bet_2;
              point_style_result = "text-red-600 font-bold"
              point_result = "-" + (totalbet * min_bet)
              break;
            case 4:
              bet_3 = min_bet;
              credit = credit - bet_3;
              point_style_result = "text-red-600 font-bold"
              point_result = "-" + (totalbet * min_bet)
              break;
          }
          
          shuffleArray_bet()
        }else{
          alert("not enough Credit")
        }
        
      }else{
        alert("Credit Empty")
      }
    }else{
      alert("Please choose minimal bet")
      flag_minimalbet = false;
    }
	};
  const call_fullbet = () => {
    flag_minimalbet = true;
    if(parseInt(min_bet) > 0){
      let credit_after = min_bet * 4;
      if(credit >= credit_after){
        count_bet = count_bet + 4
        totalbet = totalbet + 4
        bet_0 = min_bet;
        bet_1 = min_bet;
        bet_2 = min_bet;
        bet_3 = min_bet;
        sound = 0;
        spin[sound].play();
        
        credit = credit - bet_0 - bet_1 - bet_2 - bet_3;
        point_style_result = "text-red-600 font-bold"
        point_result = "-" + (totalbet * min_bet)
        shuffleArray_card(card_result_data)
        shuffleArray_fullbet()
      }else{
        alert("Credit Empty")
        flag_minimalbet = false;
      }
    }else{
      alert("Please choose minimal bet")
      flag_minimalbet = false;
    }
	};
  const call_deal = () => {
      count_bet = 4
      sound = 0;
      spin[sound].play();
      shuffleArray_deal()
	};
  function hitung(arr_id,arr_val){
    let flag_royalflush = false
    let flag_fiveofkind = false
    let flag_straight_flush = false
    let flag_fourofkind = false
    let flag_fullhouse = false
    let flag_flush = false
    let flag_straight = false
    let flag_threeofkind = false
    let flag_twopair = false
    
    flag_royalflush = royal_flush(arr_id)
    if(!flag_royalflush){
      flag_fiveofkind = five_kind(arr_id)
      if(!flag_fiveofkind){
        flag_straight_flush = straight_flush(arr_id)
        if(!flag_straight_flush){
          flag_fourofkind = fourofkind(arr_id,arr_val)
          if(!flag_fourofkind){
            flag_fullhouse = fullhouse(arr_id,arr_val)
            if(!flag_fullhouse){
              flag_flush = flush(arr_id,arr_val)
              if(!flag_flush){
                flag_straight = straight(arr_id)
                if(!flag_straight){
                  flag_threeofkind = threeofkind(arr_id,arr_val)
                  if(!flag_threeofkind){
                    flag_twopair = twopair(arr_id,arr_val)
                    if(!flag_twopair){
                      acepair(arr_id,arr_val)
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
    if(flag_win){
      sound = 0;
      win[sound].play();
    }
  }
  function royal_flush(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["val"] = temp_data.val 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].code_card]){
        counts[objdata_master[i].code_card] += 1
      }else{
        counts[objdata_master[i].code_card] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 4){
            temp.push(prop + ":" + counts[prop])
        }
    }
    if(temp.length > 0){
      let temp_string = temp[0]
      let temp_result = temp_string.split(":");
      let total_temp = temp_result[1];
      let total_jk = 0;
      let total_card = 0;
      if(parseInt(total_temp) == 5){
        for(let i=0;i<objdata_master.length;i++){
          switch(objdata_master[i].val){
            case "10":
              total_jk = total_jk + 1;break;
            case "J":
              total_jk = total_jk + 1;break;
            case "K":
              total_jk = total_jk + 1;break;
            case "Q":
              total_jk = total_jk + 1;break;
            case "AS":
              total_jk = total_jk + 1;break;
            case "JK":
              total_jk = total_jk + 1;break;
          }
        }
        total_card = total_jk
        if(total_card == 5){
          info_result = "Royal Flush"
          info_card = pattern_stright_10
          flag_win = true
          credit_animation(credit,0,totalbet)
          flag_func = true;
        }
      }
    }
    return flag_func
  }
  function five_kind(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].val_display]){
        counts[objdata_master[i].val_display] += 1
      }else{
        counts[objdata_master[i].val_display] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 3){
            temp.push(prop + ":" + counts[prop])
        }
    }
    if(temp.length > 0){
      let temp_string = temp[0]
      let temp_result = temp_string.split(":");
      let total_temp = temp_result[1];
      let total_jk = 0;
      let total_card = 0;
      if(parseInt(total_temp) == 4){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
        if(total_card == 5){
          info_result = "5 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true
          credit_animation(credit,1,totalbet)
          flag_func = true;
        }
      }
      if(parseInt(total_temp) == 3){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
        if(total_card == 5){
          info_result = "5 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true
          credit_animation(credit,1,totalbet)
          flag_func = true;
        }
      }
    }
    return flag_func
  }
  function straight_flush(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].code_card]){
        counts[objdata_master[i].code_card] += 1
      }else{
        counts[objdata_master[i].code_card] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 5){
            temp.push(prop + ":" + counts[prop])
        }
    }
    if(temp.length > 0){
      let obj_final = []
      let objdata_final = []
      let temp_string = ""
      let temp_result;
      for(let i=0;i<arr_id.length;i++){
          let temp_data = card_result_data.find(card => card.id == arr_id[i])
          temp_string = temp[0]
          temp_result = temp_string.split(":");
          if(temp_result[0] == temp_data.code_card){
            obj_final["id"] = temp_data.id 
            obj_final["code_card"] = temp_data.code_card 
            obj_final["val_display"] = temp_data.val_display 
            objdata_final.push(temp_data.val_display)
            obj_final = []
          }
      }
     
      let flag = []
      flag[0] = checkArray(pattern_stright_1,objdata_final)
      flag[1] = checkArray(pattern_stright_2,objdata_final)
      flag[2] = checkArray(pattern_stright_3,objdata_final)
      flag[3] = checkArray(pattern_stright_4,objdata_final)
      flag[4] = checkArray(pattern_stright_5,objdata_final)
      flag[5] = checkArray(pattern_stright_6,objdata_final)
      flag[6] = checkArray(pattern_stright_7,objdata_final)
      flag[7] = checkArray(pattern_stright_8,objdata_final)
      flag[8] = checkArray(pattern_stright_9,objdata_final)
      flag[9] = checkArray(pattern_stright_10,objdata_final)
      for(let i=0;i<flag.length;i++){
        if(flag[i] == true){
          info_result = "STRAIGHT FLUSH"
          info_card = pattern_stright_10
          flag_win = true
          credit_animation(credit,2,totalbet)
          flag_func = true;
          break;
        }
      }
    }
    return flag_func;
  }
  function fourofkind(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      if(counts[arr_val[i]]){
        counts[arr_val[i]] += 1
      }else{
        counts[arr_val[i]] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 4){
            temp.push(prop + ":" + counts[prop])
        }
    }
    let total = 0;
    let total_temp = temp.length
    let temp_string = ""
    let temp_result;
    for(let i=0;i<total_temp;i++){
        temp_string = temp[i]
        temp_result = temp_string.split(":");
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    if(total == 4){//FOUR OF KIND
      console.log("TOTAL :"+total+" FOUR OF KIND")
      info_result = "FOUR OF KIND"
      info_card = temp
      flag_win = true
      flag = true
      credit_animation(credit,3,totalbet)
    }
    return flag
  }
  function fullhouse(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      if(counts[arr_val[i]]){
        counts[arr_val[i]] += 1
      }else{
        counts[arr_val[i]] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 2){
            temp.push(prop + ":" + counts[prop])
        }
    }
    let total = 0;
    let total_temp = temp.length
    let temp_string = ""
    let temp_result;
    for(let i=0;i<total_temp;i++){
        temp_string = temp[i]
        temp_result = temp_string.split(":");
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    if(total == 5){//FULL HOUSE
      console.log("TOTAL :"+total+" FULL HOUSE")
      info_result = "FULL HOUSE"
      info_card = temp
      flag_win = true
      flag = true
      credit_animation(credit,4,totalbet)
    }
    return flag
  }
  function flush(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      if(counts[temp_data.code_card]){
        counts[temp_data.code_card] += 1
      }else{
        counts[temp_data.code_card] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 5){
            temp.push(prop + ":" + counts[prop])
        }
    }
    console.log(temp)
    let total = 0;
    let total_temp = temp.length
    let temp_string = ""
    let temp_result;
    for(let i=0;i<total_temp;i++){
        temp_string = temp[i]
        temp_result = temp_string.split(":");
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    if(total_temp == 1){
      if(total == 5){ //FLUSH
        console.log("TOTAL :"+total+" FLUSH")
        info_result = "FLUSH"
        info_card = temp
        flag_win = true
        flag = false
        credit_animation(credit,5,totalbet)
      }
    }
    return flag
  }
  function straight(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(temp_data.val_display)
      obj = []
    }
    
    let flag = []
    
    flag[0] = checkArray(pattern_stright_1,objdata_master)
    flag[1] = checkArray(pattern_stright_2,objdata_master)
    flag[2] = checkArray(pattern_stright_3,objdata_master)
    flag[3] = checkArray(pattern_stright_4,objdata_master)
    flag[4] = checkArray(pattern_stright_5,objdata_master)
    flag[5] = checkArray(pattern_stright_6,objdata_master)
    flag[6] = checkArray(pattern_stright_7,objdata_master)
    flag[7] = checkArray(pattern_stright_8,objdata_master)
    flag[8] = checkArray(pattern_stright_9,objdata_master)
    flag[9] = checkArray(pattern_stright_10,objdata_master)
    for(let i=0;i<flag.length;i++){
      if(flag[i] == true){
        info_result = "STRAIGHT"
        info_card = pattern_stright_10
        flag_win = true
        credit_animation(credit,6,totalbet)
        flag_func = true;
        break;
      }
    }
    return flag_func;
  }
  function threeofkind(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].val_display]){
        counts[objdata_master[i].val_display] += 1
      }else{
        counts[objdata_master[i].val_display] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 2){
            temp.push(prop + ":" + counts[prop])
        }
    }
    if(temp.length > 0){
      let temp_string = temp[0]
      let temp_result = temp_string.split(":");
      let total_temp = temp_result[1];
      let total_jk = 0;
      let total_card = 0;
      if(parseInt(total_temp) == 3){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
        if(total_card == 3){
          info_result = "3 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true
          credit_animation(credit,7,totalbet)
          flag_func = true;
        }
      }
      if(parseInt(total_temp) == 2){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
        if(total_card == 3){
          info_result = "3 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true
          credit_animation(credit,7,totalbet)
          flag_func = true;
        }
      }
    }
    return flag_func
  }
  function twopair(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      if(counts[arr_val[i]]){
        counts[arr_val[i]] += 1
      }else{
        counts[arr_val[i]] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 2){
            temp.push(prop + ":" + counts[prop])
        }
    }
    let total = 0;
    let total_temp = temp.length
    let temp_string = ""
    let temp_result;
    for(let i=0;i<total_temp;i++){
        temp_string = temp[i]
        temp_result = temp_string.split(":");
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    let flag_two = false
    if(total_temp == 2){
      if(total == 4 || total == 6){
        for(let x=0;x<temp.length;x++){
          temp_result = temp[x].split(":");
          switch(temp_result[0]){
            case "10":
              flag_two = true;break;
            case "J":
              flag_two = true;break;
            case "Q":
              flag_two = true;break;
            case "K":
              flag_two = true;break;
            case "AS":
              flag_two = true;break;
            case "JK":
              flag_two = true;break;
          }
        }
        if(flag_two){//2 PAIR
          info_result = "2 PAIR"
          info_card = temp
          flag_win = true
          flag = true
          credit_animation(credit,8,totalbet)
        }
      }
    }
    return flag
  }
  function acepair(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val"] = temp_data.val 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].val_display]){
        counts[objdata_master[i].val_display] += 1
      }else{
        counts[objdata_master[i].val_display] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 2){
            temp.push(prop + ":" + counts[prop])
        }
    }
    let total_as = 0;
    let total_jk = 0;
    let total_card = 0;
    if(temp.length > 0){
      let temp_string = temp[0]
      let temp_result = temp_string.split(":");
      let total_temp = temp_result[1];
      
      if(parseInt(total_temp) == 2){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].val == "AS"){
            total_as = total_as + 1;
          }
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        
        
        if(total_as == 2){ // 2 AS
          info_result = "ACE PAIR"
          info_card = temp
          flag_win = true
          credit_animation(credit,8,totalbet)
          flag_func = true;
        }
      }
    }else{
      for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].val == "AS"){
            total_as = total_as + 1;
          }
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
      }
      total_card = total_as + total_jk
      if(total_card == 2){ // 1 as + 1 jk
          info_result = "ACE PAIR"
          info_card = temp
          flag_win = true
          credit_animation(credit,8,totalbet)
          flag_func = true;
      }
    }
    return flag_func
  }
  function checkArray(arr_1,arr_2){
    return arr_1.every((val) => arr_2.includes(val))
  }
  
  function shuffleArray_card(array){
    let i = 0
    while(i<7){
      let randomNumber = Math.floor(Math.random() * array.length)
      if(!usedIndexes.includes(randomNumber)){
        shuffleArray.push(array[randomNumber]);
        usedIndexes.push(randomNumber);
        i++;
      }
    }
    
    // shuffleArray.push(array[12]);
    // shuffleArray.push(array[2]);
    // shuffleArray.push(array[26]);
    // shuffleArray.push(array[49]);
    // shuffleArray.push(array[20]);
    // shuffleArray.push(array[30]);
    // shuffleArray.push(array[52]);
    console.log(shuffleArray)
  }
  function shuffleArray_bet(){
    if(count_bet == 4){
      flag_bet = false
    }
    
    if(count_bet == 1){
      card_result_0_id = shuffleArray[0].id
      card_result_2_id = shuffleArray[2].id
      card_result_0_img = shuffleArray[0].img
      card_result_1_img = "./"+path_card+"CARD_FLOP.png"
      card_result_2_img = shuffleArray[2].img
      card_result_0_val = shuffleArray[0].val
      card_result_1_val = "NULL"
      card_result_2_val = shuffleArray[2].val
    }
    if(count_bet == 2){
      
      card_result_4_id = shuffleArray[4].id
      card_result_3_img = "./"+path_card+"CARD_FLOP.png"
      card_result_4_img = shuffleArray[4].img
      card_result_3_val = "NULL"
      card_result_4_val = shuffleArray[4].val
     
    }
    if(count_bet == 3){
      
      card_result_5_id = shuffleArray[5].id
      card_result_5_img = shuffleArray[5].img
      card_result_5_val = shuffleArray[5].val
    }
    if(count_bet == 4){
      card_result_1_id = shuffleArray[1].id
      card_result_3_id = shuffleArray[3].id
      card_result_6_id = shuffleArray[6].id
      card_result_1_img = shuffleArray[1].img
      card_result_3_img = shuffleArray[3].img
      card_result_6_img = shuffleArray[6].img
      card_result_1_val = shuffleArray[1].val
      card_result_3_val = shuffleArray[3].val
      card_result_6_val = shuffleArray[6].val

      card_result_array_id.push(card_result_0_id)
      card_result_array_id.push(card_result_1_id)
      card_result_array_id.push(card_result_2_id)
      card_result_array_id.push(card_result_3_id)
      card_result_array_id.push(card_result_4_id)
      card_result_array_id.push(card_result_5_id)
      card_result_array_id.push(card_result_6_id)

      card_result_array_val.push(card_result_0_val)
      card_result_array_val.push(card_result_1_val)
      card_result_array_val.push(card_result_2_val)
      card_result_array_val.push(card_result_3_val)
      card_result_array_val.push(card_result_4_val)
      card_result_array_val.push(card_result_5_val)
      card_result_array_val.push(card_result_6_val)
      // console.log(card_result_array_val)
      // console.log(card_result_array_id)
      hitung(card_result_array_id,card_result_array_val);
    }
  
    return shuffleArray;
  }
  function shuffleArray_fullbet(){
    if(count_bet == 4){
      flag_bet = false
    }
    card_result_0_id = shuffleArray[0].id
    card_result_1_id = shuffleArray[1].id
    card_result_2_id = shuffleArray[2].id
    card_result_3_id = shuffleArray[3].id
    card_result_4_id = shuffleArray[4].id
    card_result_5_id = shuffleArray[5].id
    card_result_6_id = shuffleArray[6].id
    card_result_0_img = shuffleArray[0].img
    card_result_1_img = shuffleArray[1].img
    card_result_2_img = shuffleArray[2].img
    card_result_3_img = shuffleArray[3].img
    card_result_4_img = shuffleArray[4].img
    card_result_5_img = shuffleArray[5].img
    card_result_6_img = shuffleArray[6].img
    card_result_0_val = shuffleArray[0].val
    card_result_1_val = shuffleArray[1].val
    card_result_2_val = shuffleArray[2].val
    card_result_3_val = shuffleArray[3].val
    card_result_4_val = shuffleArray[4].val
    card_result_5_val = shuffleArray[5].val
    card_result_6_val = shuffleArray[6].val

    card_result_array_id.push(card_result_0_id)
    card_result_array_id.push(card_result_1_id)
    card_result_array_id.push(card_result_2_id)
    card_result_array_id.push(card_result_3_id)
    card_result_array_id.push(card_result_4_id)
    card_result_array_id.push(card_result_5_id)
    card_result_array_id.push(card_result_6_id)

    card_result_array_val.push(card_result_0_val)
    card_result_array_val.push(card_result_1_val)
    card_result_array_val.push(card_result_2_val)
    card_result_array_val.push(card_result_3_val)
    card_result_array_val.push(card_result_4_val)
    card_result_array_val.push(card_result_5_val)
    card_result_array_val.push(card_result_6_val)
    hitung(card_result_array_id,card_result_array_val);
  
    return shuffleArray;
  }
  function shuffleArray_deal(){
    if(count_bet == 4){
      flag_bet = false
    }
    console.log("total bet : " + totalbet)
    card_result_0_id = shuffleArray[0].id
    card_result_1_id = shuffleArray[1].id
    card_result_2_id = shuffleArray[2].id
    card_result_3_id = shuffleArray[3].id
    card_result_4_id = shuffleArray[4].id
    card_result_5_id = shuffleArray[5].id
    card_result_6_id = shuffleArray[6].id
    card_result_0_img = shuffleArray[0].img
    card_result_1_img = shuffleArray[1].img
    card_result_2_img = shuffleArray[2].img
    card_result_3_img = shuffleArray[3].img
    card_result_4_img = shuffleArray[4].img
    card_result_5_img = shuffleArray[5].img
    card_result_6_img = shuffleArray[6].img
    card_result_0_val = shuffleArray[0].val
    card_result_1_val = shuffleArray[1].val
    card_result_2_val = shuffleArray[2].val
    card_result_3_val = shuffleArray[3].val
    card_result_4_val = shuffleArray[4].val
    card_result_5_val = shuffleArray[5].val
    card_result_6_val = shuffleArray[6].val

    card_result_array_id.push(card_result_0_id)
    card_result_array_id.push(card_result_1_id)
    card_result_array_id.push(card_result_2_id)
    card_result_array_id.push(card_result_3_id)
    card_result_array_id.push(card_result_4_id)
    card_result_array_id.push(card_result_5_id)
    card_result_array_id.push(card_result_6_id)

    card_result_array_val.push(card_result_0_val)
    card_result_array_val.push(card_result_1_val)
    card_result_array_val.push(card_result_2_val)
    card_result_array_val.push(card_result_3_val)
    card_result_array_val.push(card_result_4_val)
    card_result_array_val.push(card_result_5_val)
    card_result_array_val.push(card_result_6_val)
    
    hitung(card_result_array_id,card_result_array_val);
  }
  function credit_animation(credit_before,n,total_bet){
    console.log("CREDIT BEFORE : " + credit)
    console.log("POINT : " + list_point[n].poin* total_bet)
    point_result = "+" + list_point[n].poin* total_bet
    let credit_target = credit_before + (list_point[n].poin * total_bet)
    
    console.log("credit before : "+ credit)
    console.log("credit after : "+ credit_target)
    
    animateValue(credit, credit_target, 500)
    point_style_result = "text-blue-600 font-bold";
    list_point_style = "bg-slate-500"
    list_point_id = list_point[n].id
    flag_all = false
  }
  function animateValue(start, end, duration) {
    // assumes integer values for start and end
    
   
    var range = end - start;
    // no timer shorter than 50ms (not really visible any way)
    var minTimer = 50;
    // calc step time to show all interediate values
    var stepTime = Math.abs(Math.floor(duration / range));
    
    // never go below minTimer
    stepTime = Math.max(stepTime, minTimer);
    
    // get current time and calculate desired end time
    var startTime = new Date().getTime();
    var endTime = startTime + duration;
    var timer;
  
    function run() {
        var now = new Date().getTime();
        var remaining = Math.max((endTime - now) / duration, 0);
        var value = Math.round(end - (remaining * range));
        credit = value;
        if (value == end) {
            clearInterval(timer);
            flag_all = true;
        }
    }
    
    timer = setInterval(run, stepTime);
    run();
}
</script>

<main class="container mx-auto lg:px-2 text-base-content glass xl:rounded-box xl:mt-7 max-w-screen-xl bg-opacity-60 pb-5 xl:pb-5">
  <section class="grid grid-cols-1 w-full ">
    <article class="select-none mt-2 w-full ">
      <div class="p-0">
        <div class="text-xs lg:text-lg bg-base-100 rounded-md p-2">
          CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(credit)}</span>
          <span class="{point_style_result}">{point_result}</span>
        </div>
      </div>
    </article>
  </section>
  <section class="w-full select-none rounded-md p-2 mt-2 bg-base-100  grid grid-cols-2 ">
    <div class="w-1/2">
        <div class="grid grid-cols-2 ">
          <div class="text-left">BET 1</div>
          <div class="text-right text-lg link-accent">{new Intl.NumberFormat().format(bet_0)}</div>
        </div>
        <div class="grid grid-cols-2 ">
          <div class="text-left">BET 2</div>
          <div class="text-right text-lg link-accent">{new Intl.NumberFormat().format(bet_1)}</div>
        </div>
        <div class="grid grid-cols-2 ">
          <div class="text-left">BET 3</div>
          <div class="text-right text-lg link-accent">{new Intl.NumberFormat().format(bet_2)}</div>
        </div>
        <div class="grid grid-cols-2 ">
          <div class="text-left">BET 4</div>
          <div class="text-right text-lg link-accent">{new Intl.NumberFormat().format(bet_3)}</div>
        </div>
    </div>
    <div class="grid grid-cols-2">
      <div class="flex w-full pr-20 {list_point[0]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[0]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[0]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pl-10 pr-10 {list_point[5]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[5]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[5]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pr-20 {list_point[1]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[1]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[1]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pl-10 pr-10 {list_point[6]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[6]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[6]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pr-20 {list_point[2]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[2]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[2]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pl-10 pr-10 {list_point[7]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[7]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[7]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pr-20 {list_point[3]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[3]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[3]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pl-10 pr-10 {list_point[8]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[8]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[8]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pr-20 {list_point[4]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[4]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[4]["poin"]*totalbet)}</div>
      </div>
      <div class="flex w-full pl-10 pr-10 {list_point[9]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full">{list_point[9]["name"]}</div>
        <div class="w-full text-right link-accent">{new Intl.NumberFormat().format(list_point[9]["poin"]*totalbet)}</div>
      </div>
    </div>
    
  </section>
  
  
  <section class="w-full flex justify-center mt-5 ">
    <img class="" src="{card_result_0_img}" alt="">
    <img class="" src="{card_result_1_img}" alt="">
    <img class="" src="{card_result_2_img}" alt="">
    <img class="" src="{card_result_3_img}" alt="">
    <img class="" src="{card_result_4_img}" alt="">
    <img class="" src="{card_result_5_img}" alt="">
    <img class="" src="{card_result_6_img}" alt="">
  </section>
 
  <section class="flex w-full mt-5 gap-2">
    <div class="flex flex-col  w-full">
      <span class="self-end text-lg">Minimal Bet</span>
      <select
        disabled={flag_minimalbet}
        bind:value={min_bet} 
        class="select select-bordered max-w-xs text-lg self-end">
        <option value=100>100</option>
        <option value=200>200</option>
        <option value=300>300</option>
        <option value=400>400</option>
        <option value=500>500</option>
        <option value=600>600</option>
        <option value=700>700</option>
        <option value=800>800</option>
        <option value=900>900</option>
        <option value=1000>1000</option>
      </select>
    </div>
    <div class="flex place-self-end gap-1 w-full">
      {#if flag_all}
        {#if !flag_bet}
          <button on:click={() => {
              call_play();
            }} class="btn btn-success btn-md ">Play</button>
        {/if}
        {#if flag_bet}
          <button on:click={() => {
              call_bet();
            }} class="btn btn-primary btn-md " >BET</button>
          {#if flag_fullbet}
          <button on:click={() => {
              call_fullbet();
            }} class="btn btn-primary btn-md " >FULL BET</button>
          {/if}
          {#if flag_deal}
          <button on:click={() => {
              call_deal();
            }} class="btn btn-primary btn-md ">DEAL</button>
          {/if}
        {/if}
      {/if}
    </div>
  </section>
 
</main>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
