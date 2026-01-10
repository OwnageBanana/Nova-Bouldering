import sharpIcon from "../assets/icons/knife.png";
import strongIcon from "../assets/icons/muscle.png";
import delicateIcon from "../assets/icons/thin.png";
import highballIcon from "../assets/icons/mountain.png";


export function getAttributeImage(attr) {

  switch(attr) {
    case "strong":
      return strongIcon;
    case "highball":
      return highballIcon;
    case "delicate":
      return delicateIcon;
    case "sharp":
      return sharpIcon;

  }

}