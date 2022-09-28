<script>

  let credit = 5000
  let min_bet = 50
  let totalbet = 0
  let bet_0 = 0
  let bet_1 = 0
  let bet_2 = 0
  let bet_3 = 0
  let sound = 0;
  let path_card = "CARD/WHITE/";
  let spin = [
    new Audio("/sounds/spin.mp3"),
  ];
  let win = [
    new Audio("/sounds/win.mp3"),
    new Audio("/sounds/win.mp3"),
  ];
  let list_point = [
    {id:"1",code:"RF",name:"Royal Flush",poin:25000},
    {id:"2",code:"5K",name:"5 Of A Kind",poin:10000},
    {id:"3",code:"SF",name:"Straight Flush",poin:6000},
    {id:"4",code:"4K",name:"4 Of A Kind",poin:2500},
    {id:"5",code:"FH",name:"Full House",poin:350},
    {id:"6",code:"FL",name:"Flush",poin:250},
    {id:"7",code:"ST",name:"Straight",poin:150},
    {id:"8",code:"3K",name:"3 Of A Kind",poin:100},
    {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:50},
    {id:"10",code:"AP",name:"Ace Pair",poin:50},
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
  let card_background = "bg-base-100"
  let card_result_0_id = "NULL"
  let card_result_1_id = "NULL"
  let card_result_2_id = "NULL"
  let card_result_3_id = "NULL"
  let card_result_4_id = "NULL"
  let card_result_5_id = "NULL"
  let card_result_6_id = "NULL"
  let card_result_0_img = ""
  let card_result_1_img = ""
  let card_result_2_img = ""
  let card_result_3_img = ""
  let card_result_4_img = ""
  let card_result_5_img = ""
  let card_result_6_img = ""
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
    card_result_0_img = ""
    card_result_1_img = ""
    card_result_2_img = ""
    card_result_3_img = ""
    card_result_4_img = ""
    card_result_5_img = ""
    card_result_6_img = ""
    card_result_0_val = ""
    card_result_1_val = ""
    card_result_2_val = ""
    card_result_3_val = ""
    card_result_4_val = ""
    card_result_5_val = ""
    card_result_6_val = ""
	};
  const call_bet = () => {
      if(credit > 0){
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
        alert("Credit Empty")
      }
	};
  const call_fullbet = () => {
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
      }
	};
  const call_deal = () => {
      count_bet = 4
      sound = 0;
      spin[sound].play();
      shuffleArray_deal()
	};
  function hitung(arr_id,arr_val){
    // let temp_result = [];
    // temp_result = count_duplicate(arr_id,arr_val)
    // console.log(temp_result)
    let flag_fourofkind = false
    let flag_fullhouse = false
    let flag_flush = false
    let flag_straight = false
    let flag_threeofkind = false
    let flag_twopair = false
    
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
    
    if(flag_win){
      sound = 0;
      win[sound].play();
    }
  }
  function royal_flush(){}
  function five_kind(){}
  
  function straight_flush(arr_id){
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
    console.log(counts)
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
  function threeofkind(arr_id,arr_val){
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
    if(total_temp == 2){
      if(total == 3){//3 OF KIND
        console.log("TOTAL :"+total+" 3 OF KIND")
        info_result = "3 OF KIND"
        info_card = temp
        flag_win = true
        flag = true
        credit_animation(credit,7,totalbet)
      }
    }
    return flag
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
  function acepair(arr_id,arr_val){
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
    if(total_temp == 1){
      if(total == 2){ //ACE PAIR
        temp_result = temp[0].split(":");
        if(temp_result[0] == "AS"){
          console.log("TOTAL :"+total+" ACE PAIR")
          info_result = "ACE PAIR"
          info_card = temp
          flag_win = true
          flag = false
          credit_animation(credit,8,totalbet)
        }
      }
    }
    return flag
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
    // shuffleArray = [8,21,34,47,51,2,27]
    // shuffleArray = [8,21,34,47,51,2,27]
    // console.log("TAI")
    // shuffleArray.push(array[8]);
    // shuffleArray.push(array[21]);
    // shuffleArray.push(array[34]);
    // shuffleArray.push(array[47]);
    // shuffleArray.push(array[51]);
    // shuffleArray.push(array[2]);
    // shuffleArray.push(array[27]);
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
  <section class="grid grid-cols-3 w-full ">
    <article class="select-none mt-2 w-full ">
      <div class="p-2">
        <div class="text-xs lg:text-lg bg-base-100 rounded-md p-2">
          CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(credit)}</span>
          <span class="{point_style_result}">{point_result}</span>
        </div>
        <div class="bg-base-100 rounded-md p-2 mt-2">
          <div class="text-lg link-accent">{bet_0}</div>
          <div class="text-lg link-accent">{bet_1}</div>
          <div class="text-lg link-accent">{bet_2}</div>
          <div class="text-lg link-accent">{bet_3}</div>
        </div>
        
      </div>
    </article>
    <article class="select-none mt-2 w-full ">
      <div class="flex flex-col items-center hidden">
        <div class="card bg-base-100 shadow-lg  rounded-md select-none w-1/2">
          <div class="card-body grass items-center text-center">
            <div class="text-sm h-16 mt-2">
              <center>
                <span class="text-2xl font-bold">WINNER</span>
                <br>
                <span class="text-lg  font-bold link-accent">{info_result}</span>
              </center>
            </div>
            <div class="text-sm h-16">
              <center>
                <span class="text-2xl font-bold">CARD</span>
                <br>
                <span class="text-lg  font-bold link-accent">{info_card}</span>
              </center>
            </div>
          </div>
        </div>
      </div>
      
      
    </article>
    <article class="w-full select-none rounded-md p-2 mt-2 bg-base-100 rounded-md">
      <table class="w-full">
        {#each list_point as rec}
          {#if rec.id == list_point_id}
          <tr class="{list_point_style} ">
            <td>{rec.name}</td>
            <td class="text-xs lg:text-lg link-accent text-right">{new Intl.NumberFormat().format(rec.poin*totalbet)}</td>
          </tr>
          {:else}
          <tr>
            <td>{rec.name}</td>
            <td class="text-xs lg:text-lg link-accent text-right">{new Intl.NumberFormat().format(rec.poin*totalbet)}</td>
          </tr>
          {/if}
        {/each}
      </table>
    </article>
  </section>
  <section class="grid grid-cols-7 w-full mt-2 mb-2 select-none gap-2">
    <div class="card {card_background} shadow-lg  rounded-md select-none h-[220px]">
      <img width="170" src="{card_result_0_img}" alt="">
    </div>
    <div class="card {card_background} shadow-lg  rounded-md select-none h-[220px]">
      <img width="170" src="{card_result_1_img}" alt="">
    </div>
    <div class="card {card_background} shadow-lg  rounded-md select-none h-[220px]">
      <img width="170" src="{card_result_2_img}" alt="">
    </div>
    <div class="card {card_background} shadow-lg  rounded-md select-none h-[220px]">
      <img width="170" src="{card_result_3_img}" alt="">
    </div>
    <div class="card {card_background} shadow-lg  rounded-md select-none h-[220px]">
      <img width="170" src="{card_result_4_img}" alt="">
    </div>
    <div class="card {card_background} shadow-lg  rounded-md select-none h-[220px]">
      <img width="170" src="{card_result_5_img}" alt="">
    </div>
    <div class="card {card_background} shadow-lg  rounded-md select-none h-[220px]">
      <img width="170" src="{card_result_6_img}" alt="">
    </div>
    
    
    
    
    
    
    
  </section>
  
  <center class="mt-2 mb-2">
    {#if flag_all}
      {#if !flag_bet}
        <button on:click={() => {
            call_play();
          }} class="btn btn-success btn-lg">Play</button>
      {/if}
      {#if flag_bet}
        <button on:click={() => {
            call_bet();
          }} class="btn btn-active btn-lg" >BET</button>
        {#if flag_fullbet}
        <button on:click={() => {
            call_fullbet();
          }} class="btn btn-active btn-lg" >FULL BET</button>
        {/if}
        {#if flag_deal}
        <button on:click={() => {
            call_deal();
          }} class="btn btn-active btn-lg">DEAL</button>
        {/if}
      {/if}
    {/if}
  </center>
</main>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
