<script>

  let credit = 5000
  let min_bet = 50
  let bet_0 = 0
  let bet_1 = 0
  let bet_2 = 0
  let bet_3 = 0
  let list_point = [
    {id:"1",name:"Royal Flush",poin:25000},
    {id:"2",name:"5 Of A Kind",poin:10000},
    {id:"3",name:"Straight Flush",poin:6000},
    {id:"4",name:"4 Of A Kind",poin:2500},
    {id:"5",name:"Full House",poin:350},
    {id:"6",name:"Full",poin:250},
    {id:"7",name:"Straight",poin:150},
    {id:"8",name:"3 Of A Kind",poin:100},
    {id:"9",name:"2 Pair (10 PAIR)",poin:50},
    {id:"10",name:"Ace Pair",poin:50},
  ]
  const card_result_dua = [
    {id:"2_diamond",val:"2",img:"./CARD_RED_DIAMOND_2.png"},
    {id:"3_diamond",val:"3",img:"./CARD_RED_DIAMOND_3.png"},
    {id:"4_diamond",val:"4",img:"./CARD_RED_DIAMOND_4.png"},
    {id:"5_diamond",val:"5",img:"./CARD_RED_DIAMOND_5.png"},
    {id:"6_diamond",val:"6",img:"./CARD_RED_DIAMOND_6.png"},
    {id:"7_diamond",val:"7",img:"./CARD_RED_DIAMOND_7.png"},
    {id:"8_diamond",val:"8",img:"./CARD_RED_DIAMOND_8.png"},
    {id:"9_diamond",val:"9",img:"./CARD_RED_DIAMOND_9.png"},
    {id:"10_diamond",val:"10",img:"./CARD_RED_DIAMOND_10.png"},
    {id:"j_diamond",val:"J",img:"./CARD_RED_DIAMOND_J.png"},
    {id:"q_diamond",val:"Q",img:"./CARD_RED_DIAMOND_Q.png"},
    {id:"k_diamond",val:"K",img:"./CARD_RED_DIAMOND_K.png"},
    {id:"as_diamond",val:"AS",img:"./CARD_RED_DIAMOND_AS.png"},
    {id:"jk_diamond",val:"JK",img:"./CARD_RED_DIAMOND_JK.png"},
    {id:"2_heart",val:"2",img:"./CARD_RED_LOVE_2.png"},
    {id:"3_heart",val:"3",img:"./CARD_RED_LOVE_3.png"},
    {id:"4_heart",val:"4",img:"./CARD_RED_LOVE_4.png"},
    {id:"5_heart",val:"5",img:"./CARD_RED_LOVE_5.png"},
    {id:"6_heart",val:"6",img:"./CARD_RED_LOVE_6.png"},
    {id:"7_heart",val:"7",img:"./CARD_RED_LOVE_7.png"},
    {id:"8_heart",val:"8",img:"./CARD_RED_LOVE_8.png"},
    {id:"9_heart",val:"9",img:"./CARD_RED_LOVE_9.png"},
    {id:"10_heart",val:"10",img:"./CARD_RED_LOVE_10.png"},
    {id:"j_heart",val:"J",img:"./CARD_RED_LOVE_J.png"},
    {id:"q_heart",val:"Q",img:"./CARD_RED_LOVE_Q.png"},
    {id:"k_heart",val:"K",img:"./CARD_RED_LOVE_K.png"},
    {id:"as_heart",val:"AS",img:"./CARD_RED_LOVE_AS.png"},
    {id:"jk_heart",val:"JK",img:"./CARD_RED_LOVE_JK.png"},

    {id:"2_club",val:"2",img:"./CARD_BLACK_KRITING_2.png"},
    {id:"3_club",val:"3",img:"./CARD_BLACK_KRITING_3.png"},
    {id:"4_club",val:"4",img:"./CARD_BLACK_KRITING_4.png"},
    {id:"5_club",val:"5",img:"./CARD_BLACK_KRITING_5.png"},
    {id:"6_club",val:"6",img:"./CARD_BLACK_KRITING_6.png"},
    {id:"7_club",val:"7",img:"./CARD_BLACK_KRITING_7.png"},
    {id:"8_club",val:"8",img:"./CARD_BLACK_KRITING_8.png"},
    {id:"9_club",val:"9",img:"./CARD_BLACK_KRITING_9.png"},
    {id:"10_club",val:"10",img:"./CARD_BLACK_KRITING_10.png"},
    {id:"j_club",val:"J",img:"./CARD_BLACK_KRITING_J.png"},
    {id:"q_club",val:"Q",img:"./CARD_BLACK_KRITING_Q.png"},
    {id:"k_club",val:"K",img:"./CARD_BLACK_KRITING_K.png"},
    {id:"as_club",val:"AS",img:"./CARD_BLACK_KRITING_AS.png"},
    {id:"jk_club",val:"JK",img:"./CARD_BLACK_KRITING_JK.png"},

    {id:"2_spade",val:"2",img:"./CARD_BLACK_DAUN_2.png"},
    {id:"3_spade",val:"3",img:"./CARD_BLACK_DAUN_3.png"},
    {id:"4_spade",val:"4",img:"./CARD_BLACK_DAUN_4.png"},
    {id:"5_spade",val:"5",img:"./CARD_BLACK_DAUN_5.png"},
    {id:"6_spade",val:"6",img:"./CARD_BLACK_DAUN_6.png"},
    {id:"7_spade",val:"7",img:"./CARD_BLACK_DAUN_7.png"},
    {id:"8_spade",val:"8",img:"./CARD_BLACK_DAUN_8.png"},
    {id:"9_spade",val:"9",img:"./CARD_BLACK_DAUN_9.png"},
    {id:"10_spade",val:"10",img:"./CARD_BLACK_DAUN_10.png"},
    {id:"j_spade",val:"J",img:"./CARD_BLACK_DAUN_J.png"},
    {id:"q_spade",val:"Q",img:"./CARD_BLACK_DAUN_Q.png"},
    {id:"k_spade",val:"K",img:"./CARD_BLACK_DAUN_K.png"},
    {id:"as_spade",val:"AS",img:"./CARD_BLACK_DAUN_AS.png"},
    {id:"jk_spade",val:"JK",img:"./CARD_BLACK_DAUN_JK.png"},
  ]
  let flag_bet = true
  let count_bet = 0;
  let card_result_0_id = "NULL"
  let card_result_1_id = "NULL"
  let card_result_2_id = "NULL"
  let card_result_3_id = "NULL"
  let card_result_4_id = "NULL"
  let card_result_5_id = "NULL"
  let card_result_6_id = "NULL"
  let card_result_0_img = "./CARD_BLANK.png"
  let card_result_1_img = "./CARD_BLANK.png"
  let card_result_2_img = "./CARD_BLANK.png"
  let card_result_3_img = "./CARD_BLANK.png"
  let card_result_4_img = "./CARD_BLANK.png"
  let card_result_5_img = "./CARD_BLANK.png"
  let card_result_6_img = "./CARD_BLANK.png"
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
  const call_play = () => {
    bet_0 = 0
    bet_1 = 0
    bet_2 = 0
    bet_3 = 0
    count_bet = 0;
    flag_bet = true;
    card_result_array_id = []
    card_result_array_val = []
    info_result = "";
    info_card = [];
    card_result_0_img = "./CARD_BLANK.png"
    card_result_1_img = "./CARD_BLANK.png"
    card_result_2_img = "./CARD_BLANK.png"
    card_result_3_img = "./CARD_BLANK.png"
    card_result_4_img = "./CARD_BLANK.png"
    card_result_5_img = "./CARD_BLANK.png"
    card_result_6_img = "./CARD_BLANK.png"
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
      shuffleArray(card_result_dua)
	};
  const call_deal = () => {
      count_bet = count_bet + 1
      shuffleArray_deal(card_result_dua)
	};
  function hitung(arr_id,arr_val){
    let temp_result = [];
    temp_result = count_duplicate(arr_id,arr_val)
    console.log(temp_result)
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
          console.log("TOTAL :"+total+" ACE PAIR")
          info_result = "ACE PAIR"
          info_card = temp
        }
      }
      if(total == 3){
        console.log("TOTAL :"+total+" 3 OF KIND")
        info_result = "3 OF KIND"
        info_card = temp
      }
      if(total == 4){
        console.log("TOTAL :"+total+" 4 OF KIND")
        info_result = "4 OF KIND"
        info_card = temp
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
          console.log("TOTAL :"+total+" 2 PAIR")
          info_result = "2 PAIR"
          info_card = temp
        }
      }
      if(total == 3){
        console.log("TOTAL :"+total+" 3 OF KIND")
        info_result = "3 OF KIND"
        info_card = temp
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
      bet_0 = min_bet
      card_result_0_id = shuffleArray[0].id
      card_result_2_id = shuffleArray[2].id
      card_result_0_img = shuffleArray[0].img
      card_result_1_img = "./CARD_FLOP.png"
      card_result_2_img = shuffleArray[2].img
      card_result_0_val = shuffleArray[0].val
      card_result_1_val = "NULL"
      card_result_2_val = shuffleArray[2].val
    }
    if(count_bet == 2){
      bet_1 = min_bet
      card_result_4_id = shuffleArray[4].id
      card_result_3_img = "./CARD_FLOP.png"
      card_result_4_img = shuffleArray[4].img
      card_result_3_val = "NULL"
      card_result_4_val = shuffleArray[4].val
     
    }
    if(count_bet == 3){
      bet_2 = min_bet
      card_result_5_id = shuffleArray[5].id
      card_result_5_img = shuffleArray[5].img
      card_result_5_val = shuffleArray[5].val
    }
    if(count_bet == 4){
      bet_3 = min_bet
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
      bet_0 = min_bet
      bet_1 = min_bet
      bet_2 = min_bet
      bet_3 = min_bet
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
      bet_1 = min_bet
      bet_2 = min_bet
      bet_3 = min_bet
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
      bet_2 = min_bet
      bet_3 = min_bet
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
      bet_3 = min_bet
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
    <article class="select-none mt-5 w-full ">
      <div class="p-2">
        <div class="text-lg">
          <span class="text-lg font-bold">CREDIT</span> : <span class="text-lg text-blue-800">{new Intl.NumberFormat().format(credit)}</span>
        </div>
        <div class="text-lg"><span class="text-lg text-blue-800">{bet_0}</span></div>
        <div class="text-lg"><span class="text-lg text-blue-800">{bet_1}</span></div>
        <div class="text-lg"><span class="text-lg text-blue-800">{bet_2}</span></div>
        <div class="text-lg"><span class="text-lg text-blue-800">{bet_3}</span></div>
      </div>
    </article>
    <article class="select-none mt-5 w-full ">
      <div class="flex flex-col items-center">
        <div class="flex flex-col rounded-md bg-slate-300 w-2/3 mt-20">
          <div class="text-sm h-16 mt-2">
            <center>
              <span class="text-2xl font-bold">WINNER</span>
              <br>
              <span class="text-lg text-orange-600 font-bold">{info_result}</span>
            </center>
          </div>
          <div class="text-sm h-16">
            <center>
              <span class="text-2xl font-bold">CARD</span>
              <br>
              <span class="text-lg text-orange-600 font-bold">{info_card}</span>
            </center>
          </div>
        </div>
      </div>
      
      
    </article>
    <article class="w-full bg-neutral-content select-none rounded-md p-2 mt-5">
      <table class="w-full">
        {#each list_point as rec}
          <tr>
            <td>{rec.name}</td>
            <td class="text-lg text-blue-800 text-right">{new Intl.NumberFormat().format(rec.poin)}</td>
          </tr>
        {/each}
      </table>
    </article>
  </section>
  <section class="grid grid-cols-7 w-full mt-2 mb-2 select-none">
    <img width="170" src="{card_result_0_img}" alt="">
    <img width="170" src="{card_result_1_img}" alt="">
    <img width="170" src="{card_result_2_img}" alt="">
    <img width="170" src="{card_result_3_img}" alt="">
    <img width="170" src="{card_result_4_img}" alt="">
    <img width="170" src="{card_result_5_img}" alt="">
    <img width="170" src="{card_result_6_img}" alt="">
  </section>
  
  <center class="mt-2 mb-2">
    {#if !flag_bet}
      <button on:click={() => {
          call_play();
        }} class="btn btn-secondary">Play</button>
    {/if}
    {#if flag_bet}
      <button on:click={() => {
          call_bet();
        }} class="btn btn-secondary" >BET</button>
      <button on:click={() => {
          call_deal();
        }} class="btn">DEAL</button>
    {/if}
  </center>
</main>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
