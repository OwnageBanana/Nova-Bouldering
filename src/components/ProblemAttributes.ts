import sharpIcon from "../assets/icons/knife.png";
import strongIcon from "../assets/icons/muscle.png";
import delicateIcon from "../assets/icons/thin.png";
import highballIcon from "../assets/icons/mountain.png";

export enum Attribute {
  STRONG = 'strong',
  HIGHBALL = 'highball',
  DELICATE = 'delicate',
  SHARP = 'sharp'
}

export function getAttributeImage(attr: Attribute) {

  switch(attr) {
    case Attribute.STRONG:
      return strongIcon;
    case Attribute.HIGHBALL:
      return highballIcon;
    case Attribute.DELICATE:
      return delicateIcon;
    case Attribute.SHARP:
      return sharpIcon;

  }

}