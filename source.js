
// import {Matrix} from './matrix.js';





function showMatrix(){

    const form = document.getElementById('set-matrix-size');
    const numRows = form.elements['rows'].value;
    const numCols = form.elements['cols'].value;

    let mat = new Matrix(numRows, numCols);
    console.log(mat.getRows());
    console.log(mat.getCols());

    // console.log(numRows.value)
    // console.log(numCols.value)

    // const element = document.getElementById("div1");
    // element.appendChild(para)

    function insertMatrix(){
        for(let i = 0; i < mat.getRows(); ++i){
            console.log(i)
            document.getElementById("set-matrix-values").innerHTML = 
            '<div class="message">Add Message<br>Title: <input type="text" id="title"><br>Text: <input type="text" id="message"><br><br></div>';
        }
        
    };
    insertMatrix()


    
}








class Matrix{
    constructor(r,c){
        this.rows = r;
        this.cols = c;
        this.mat = [];
    }

    getRows(){
        return this.rows;
    }
    
    getCols(){
        return this.cols;
    }

    //initialize(number[][]): void
    initialize(dubArr){
        for(let r = 0; r < this.height; ++r){
            let arr = [];
            for(let c = 0; c < this.width; ++c){
                arr.push(dubArr[r][c]);
            }
            this.mat.push(arr);
        }
    }

    //getVals(): number[][]
    getVals(){
        return this.mat;
    }

    //convertToEF(): Matrix
    convertToEF(row, col){ //use recursion
        //start at first row
        //find row with pivot in desired col
        //swap if needed (if no pivots, move on to next row)
        //make zeros all the way down
        //move on to next row

    }

    //convertToREF(): Matrix
    convertToREF(){
        //first convert to EF
    }

    //findDet(): number

    //findInverse(): Matrix

    //canMultiply(): boolean

    //multiplyMats(): Matrix

    //findB(): Matrix

    //findX(): Matrix
}
