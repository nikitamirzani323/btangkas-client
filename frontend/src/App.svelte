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
  ];
  let list_point = [
    {id:"1",code:"RF",name:"Royal Flush",poin:25000},
    {id:"2",code:"5K",name:"5 Of A Kind",poin:10000},
    {id:"3",code:"SF",name:"Straight Flush",poin:6000},
    {id:"4",code:"4K",name:"4 Of A Kind",poin:2500},
    {id:"5",code:"FH",name:"Full House",poin:350},
    {id:"6",code:"FL",name:"Full",poin:250},
    {id:"7",code:"ST",name:"Straight",poin:150},
    {id:"8",code:"3K",name:"3 Of A Kind",poin:100},
    {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:50},
    {id:"10",code:"AP",name:"Ace Pair",poin:50},
  ]
  const card_result_dua = [
    {id:"2_diamond",val:"2",img:"./CARD/WHITE/CARD_RED_DIAMOND_2.png"},
    {id:"3_diamond",val:"3",img:"./CARD/WHITE/CARD_RED_DIAMOND_3.png"},
    {id:"4_diamond",val:"4",img:"./CARD/WHITE/CARD_RED_DIAMOND_4.png"},
    {id:"5_diamond",val:"5",img:"./CARD/WHITE/CARD_RED_DIAMOND_5.png"},
    {id:"6_diamond",val:"6",img:"./CARD/WHITE/CARD_RED_DIAMOND_6.png"},
    {id:"7_diamond",val:"7",img:"./CARD/WHITE/CARD_RED_DIAMOND_7.png"},
    {id:"8_diamond",val:"8",img:"./CARD/WHITE/CARD_RED_DIAMOND_8.png"},
    {id:"9_diamond",val:"9",img:"./CARD/WHITE/CARD_RED_DIAMOND_9.png"},
    {id:"10_diamond",val:"10",img:"./CARD/WHITE/CARD_RED_DIAMOND_10.png"},
    {id:"j_diamond",val:"J",img:"./CARD/WHITE/CARD_RED_DIAMOND_J.png"},
    {id:"q_diamond",val:"Q",img:"./CARD/WHITE/CARD_RED_DIAMOND_Q.png"},
    {id:"k_diamond",val:"K",img:"./CARD/WHITE/CARD_RED_DIAMOND_K.png"},
    {id:"as_diamond",val:"AS",img:"./CARD/WHITE/CARD_RED_DIAMOND_AS.png"},
    {id:"2_heart",val:"2",img:"./CARD/WHITE/CARD_RED_LOVE_2.png"},
    {id:"3_heart",val:"3",img:"./CARD/WHITE/CARD_RED_LOVE_3.png"},
    {id:"4_heart",val:"4",img:"./CARD/WHITE/CARD_RED_LOVE_4.png"},
    {id:"5_heart",val:"5",img:"./CARD/WHITE/CARD_RED_LOVE_5.png"},
    {id:"6_heart",val:"6",img:"./CARD/WHITE/CARD_RED_LOVE_6.png"},
    {id:"7_heart",val:"7",img:"./CARD/WHITE/CARD_RED_LOVE_7.png"},
    {id:"8_heart",val:"8",img:"./CARD/WHITE/CARD_RED_LOVE_8.png"},
    {id:"9_heart",val:"9",img:"./CARD/WHITE/CARD_RED_LOVE_9.png"},
    {id:"10_heart",val:"10",img:"./CARD/WHITE/CARD_RED_LOVE_10.png"},
    {id:"j_heart",val:"J",img:"./CARD/WHITE/CARD_RED_LOVE_J.png"},
    {id:"q_heart",val:"Q",img:"./CARD/WHITE/CARD_RED_LOVE_Q.png"},
    {id:"k_heart",val:"K",img:"./CARD/WHITE/CARD_RED_LOVE_K.png"},
    {id:"as_heart",val:"AS",img:"./CARD/WHITE/CARD_RED_LOVE_AS.png"},

    {id:"2_club",val:"2",img:"./CARD/WHITE/CARD_BLACK_KRITING_2.png"},
    {id:"3_club",val:"3",img:"./CARD/WHITE/CARD_BLACK_KRITING_3.png"},
    {id:"4_club",val:"4",img:"./CARD/WHITE/CARD_BLACK_KRITING_4.png"},
    {id:"5_club",val:"5",img:"./CARD/WHITE/CARD_BLACK_KRITING_5.png"},
    {id:"6_club",val:"6",img:"./CARD/WHITE/CARD_BLACK_KRITING_6.png"},
    {id:"7_club",val:"7",img:"./CARD/WHITE/CARD_BLACK_KRITING_7.png"},
    {id:"8_club",val:"8",img:"./CARD/WHITE/CARD_BLACK_KRITING_8.png"},
    {id:"9_club",val:"9",img:"./CARD/WHITE/CARD_BLACK_KRITING_9.png"},
    {id:"10_club",val:"10",img:"./CARD/WHITE/CARD_BLACK_KRITING_10.png"},
    {id:"j_club",val:"J",img:"./CARD/WHITE/CARD_BLACK_KRITING_J.png"},
    {id:"q_club",val:"Q",img:"./CARD/WHITE/CARD_BLACK_KRITING_Q.png"},
    {id:"k_club",val:"K",img:"./CARD/WHITE/CARD_BLACK_KRITING_K.png"},
    {id:"as_club",val:"AS",img:"./CARD/WHITE/CARD_BLACK_KRITING_AS.png"},

    {id:"2_spade",val:"2",img:"./CARD/WHITE/CARD_BLACK_DAUN_2.png"},
    {id:"3_spade",val:"3",img:"./CARD/WHITE/CARD_BLACK_DAUN_3.png"},
    {id:"4_spade",val:"4",img:"./CARD/WHITE/CARD_BLACK_DAUN_4.png"},
    {id:"5_spade",val:"5",img:"./CARD/WHITE/CARD_BLACK_DAUN_5.png"},
    {id:"6_spade",val:"6",img:"./CARD/WHITE/CARD_BLACK_DAUN_6.png"},
    {id:"7_spade",val:"7",img:"./CARD/WHITE/CARD_BLACK_DAUN_7.png"},
    {id:"8_spade",val:"8",img:"./CARD/WHITE/CARD_BLACK_DAUN_8.png"},
    {id:"9_spade",val:"9",img:"./CARD/WHITE/CARD_BLACK_DAUN_9.png"},
    {id:"10_spade",val:"10",img:"./CARD/WHITE/CARD_BLACK_DAUN_10.png"},
    {id:"j_spade",val:"J",img:"./CARD/WHITE/CARD_BLACK_DAUN_J.png"},
    {id:"q_spade",val:"Q",img:"./CARD/WHITE/CARD_BLACK_DAUN_Q.png"},
    {id:"k_spade",val:"K",img:"./CARD/WHITE/CARD_BLACK_DAUN_K.png"},
    {id:"as_spade",val:"AS",img:"./CARD/WHITE/CARD_BLACK_DAUN_AS.png"},
    {id:"jk_black",val:"JK",img:"./CARD/WHITE/CARD_JOKER_BLACK.png"},
    {id:"jk_red",val:"JK",img:"./CARD/WHITE/CARD_JOKER_RED.png"},
  ]
  let flag_bet = true
  let flag_fullbet = true
  let flag_deal = false
  let count_bet = 0;
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
    card_result_array_id = []
    card_result_array_val = []
    info_result = "";
    info_card = [];
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
      count_bet = count_bet + 1
      totalbet = totalbet + 1
      sound = 0;
      spin[sound].play();
      switch(count_bet){
        case 1:
          flag_deal = true;
          flag_fullbet = false;
          bet_0 = min_bet;
          credit = credit - bet_0;break;
        case 2:
          bet_1 = min_bet;
          credit = credit - bet_1;break;
        case 3:
          bet_2 = min_bet;
          credit = credit - bet_2;break;
        case 4:
          bet_3 = min_bet;
          credit = credit - bet_3;break;
      }
      shuffleArray(card_result_dua)
	};
  const call_fullbet = () => {
      count_bet = count_bet + 4
      totalbet = totalbet + 4
      bet_0 = min_bet;
      bet_1 = min_bet;
      bet_2 = min_bet;
      bet_3 = min_bet;
      
      sound = 0;
      spin[sound].play();
      
      credit = credit - bet_0 - bet_1 - bet_2 - bet_3;
      shuffleArray_fullbet(card_result_dua)
	};
  const call_deal = () => {
      count_bet = count_bet + 1
      sound = 0;
      spin[sound].play();
      shuffleArray_deal(card_result_dua)
	};
  function hitung(arr_id,arr_val){
    let temp_result = [];
    temp_result = count_duplicate(arr_id,arr_val)
    console.log(temp_result)
    if(flag_win){
      sound = 0;
      win[sound].play();
    }
  }
  function count_duplicate(arr_id,arr_val){
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
          console.log(list_point[9].poin)
          console.log("TOTAL :"+total+" ACE PAIR")
          info_result = "ACE PAIR"
          info_card = temp
          flag_win = true
        }
      }
      if(total == 3){
        console.log("TOTAL :"+total+" 3 OF KIND")
        info_result = "3 OF KIND"
        info_card = temp
        flag_win = true
      }
      if(total == 4){
        console.log("TOTAL :"+total+" 4 OF KIND")
        info_result = "4 OF KIND"
        info_card = temp
        flag_win = true
      }
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
        if(flag_two){
          console.log(list_point[8].poin)
          console.log("TOTAL :"+total+" 2 PAIR")
          info_result = "2 PAIR"
          info_card = temp
          flag_win = true
        }
      }
      if(total == 3){
        console.log("TOTAL :"+total+" 3 OF KIND")
        info_result = "3 OF KIND"
        info_card = temp
        flag_win = true
      }
    }
    
    
    if(total == 5){
      console.log("TOTAL :"+total+" FULL HOUSE")
      info_result = "FULL HOUSE"
      info_card = temp
    }
    return temp
    
  }
  function shuffleArray(array){
    let shuffleArray = [];
    let usedIndexes = [];
    let i = 0
    
    if(count_bet == 4){
      flag_bet = false
    }
    while(i<7){
      let randomNumber = Math.floor(Math.random() * array.length)
      if(!usedIndexes.includes(randomNumber)){
        shuffleArray.push(array[randomNumber]);
        usedIndexes.push(randomNumber);
        i++;
      }
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
      console.log(card_result_array_val)
      console.log(card_result_array_id)
      hitung(card_result_array_id,card_result_array_val);
    }
  
    return shuffleArray;
  }
  function shuffleArray_fullbet(array){
    let shuffleArray = [];
    let usedIndexes = [];
    let i = 0

    if(count_bet == 4){
      flag_bet = false
      
    }
    
    while(i<7){
      let randomNumber = Math.floor(Math.random() * array.length)
      if(!usedIndexes.includes(randomNumber)){
        shuffleArray.push(array[randomNumber]);
        usedIndexes.push(randomNumber);
        i++;
      }
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
    console.log(card_result_array_val)
    console.log(card_result_array_id)
    hitung(card_result_array_id,card_result_array_val);
  
    return shuffleArray;
  }
  function shuffleArray_deal(array){
    let shuffleArray = [];
    let usedIndexes = [];
    let i = 0
    let n = 7
    if(count_bet == 1){
      n = 7
      flag_bet = false
    }
    if(count_bet == 2){
      n = 5
      flag_bet = false
    }
    if(count_bet == 3){
      n = 4
      flag_bet = false
    }
    if(count_bet == 4){
      n = 3
      flag_bet = false
    }
    while(i<n){
      let randomNumber = Math.floor(Math.random() * array.length)
      if(!usedIndexes.includes(randomNumber)){
        shuffleArray.push(array[randomNumber]);
        usedIndexes.push(randomNumber);
        i++;
      }
    }
    console.log(shuffleArray)
    if(count_bet == 1){
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
      console.log(card_result_array_val)
      console.log(card_result_array_id)
      hitung(card_result_array_id,card_result_array_val);
    }
    if(count_bet == 2){
      card_result_1_id = shuffleArray[0].id
      card_result_3_id = shuffleArray[1].id
      card_result_4_id = shuffleArray[2].id
      card_result_5_id = shuffleArray[3].id
      card_result_6_id = shuffleArray[4].id
      card_result_1_img= shuffleArray[0].img
      card_result_3_img = shuffleArray[1].img
      card_result_4_img = shuffleArray[2].img
      card_result_5_img = shuffleArray[3].img
      card_result_6_img = shuffleArray[4].img
      card_result_1_val = shuffleArray[0].val
      card_result_3_val = shuffleArray[1].val
      card_result_4_val = shuffleArray[2].val
      card_result_5_val = shuffleArray[3].val
      card_result_6_val = shuffleArray[4].val

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
      console.log(card_result_array_val)
      console.log(card_result_array_id)
      hitung(card_result_array_id,card_result_array_val);
    }
    if(count_bet == 3){
      card_result_1_id = shuffleArray[0].id
      card_result_3_id = shuffleArray[1].id
      card_result_5_id = shuffleArray[2].id
      card_result_6_id = shuffleArray[3].id
      card_result_1_img= shuffleArray[0].img
      card_result_3_img = shuffleArray[1].img
      card_result_5_img = shuffleArray[2].img
      card_result_6_img = shuffleArray[3].img
      card_result_1_val = shuffleArray[0].val
      card_result_3_val = shuffleArray[1].val
      card_result_5_val = shuffleArray[2].val
      card_result_6_val = shuffleArray[3].val

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
      console.log(card_result_array_val)
      console.log(card_result_array_id)
      hitung(card_result_array_id,card_result_array_val);
    }
    if(count_bet == 4){
      card_result_1_id = shuffleArray[0].id
      card_result_3_id = shuffleArray[1].id
      card_result_6_id = shuffleArray[2].id
      card_result_1_img= shuffleArray[0].img
      card_result_3_img = shuffleArray[1].img
      card_result_6_img = shuffleArray[2].img
      card_result_1_val = shuffleArray[0].val
      card_result_3_val = shuffleArray[1].val
      card_result_6_val = shuffleArray[2].val

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
      console.log(card_result_array_val)
      console.log(card_result_array_id)
      hitung(card_result_array_id,card_result_array_val);
    }
    return shuffleArray;
  }
</script>

<main class="container mx-auto lg:px-2 text-base-content glass xl:rounded-box xl:mt-7 max-w-screen-xl bg-opacity-60 pb-5 xl:pb-5">
  <section class="grid grid-cols-3 w-full ">
    <article class="select-none mt-2 w-full ">
      <div class="p-2">
        <div class="text-xs lg:text-lg">
          CREDIT : IDR <span class="link-accent">{new Intl.NumberFormat().format(credit)}</span>
        </div>
        <div class="text-lg link-accent">{bet_0}</div>
        <div class="text-lg link-accent">{bet_1}</div>
        <div class="text-lg link-accent">{bet_2}</div>
        <div class="text-lg link-accent">{bet_3}</div>
      </div>
    </article>
    <article class="select-none mt-2 w-full ">
      <div class="flex flex-col items-center">
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
    <article class="w-full select-none rounded-md p-2 mt-2">
      <table class="w-full">
        {#each list_point as rec}
          <tr>
            <td>{rec.name}</td>
            <td class="text-xs lg:text-lg link-accent text-right">{new Intl.NumberFormat().format(rec.poin*totalbet)}</td>
          </tr>
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
  </center>
</main>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
