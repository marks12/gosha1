package cmd

const commonContent = `

export function findItemIndex(list, fn) {
    let index;
    
    if (!list || !fn) return;
    
    for (let i = 0; i < list.length; i++) {
        
        if (fn(list[i])) {
            index = i;
            break;
        }
        
    }
    return index;
}



`