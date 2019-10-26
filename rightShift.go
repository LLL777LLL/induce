package rightShift

/*
����Ӧ�ã�
��ʱ������Ŀ�п��ܻ��õ�ͬ���Ĳ�ͬ���࣬����ͬ���ඨ��һ��Ψһʶ���ֶΣ����磺
�� << ��������Ψһʶ��ţ�1,2,4��8,16 ����������֮������ǲ�ͬ����ĺ� �� �磺1+4+8 = 13
���� �õ� 13 ����Ҫ�������� ��������Щ�ࡣ
*/

import (
	"math"
)

func Right(num int)[]int {
	var (
		arr1 []int
		arr []int
		count int
	)

	//��ʮ����ת���ɶ����ƣ���arr1 ��Ƭ�洢�������Ƕ����ƽ���ĵ��򣬱���֮���� range �� k ����ת���� ��������
	for ;num > 0;num /= 2 {
		count = num % 2
		arr1 = append(arr1,count)
	}

	//��Ϊ�洢��ʱ���Ƕ����ƽ������洢�ģ����Ա������������õ� 2 ����
	for k,v := range arr1 {
		if v == 1 {
			numb := math.Pow(2,float64(k))
			arr = append(arr,int(numb))
		}
	}
	return arr
}

func HasRight(num1,num2 int) bool {
	arr := Right(num2)
	for _,v := range arr {
		if v == num1 {
			return true
		}
	}
	return false
}