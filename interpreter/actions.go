package interpreter

import (
	"VenariBot/requests/expeditions"
	"VenariBot/requests/expeditions/battle"
	"VenariBot/requests/inventory"
	"VenariBot/requests/user"
	"VenariBot/strats"
	"errors"
	"strconv"
	"strings"
)

var (
	ActionList = map[string]func(*InterpreterTask, *strats.VenariTask, []string) error {
		"getUserEnergy": getUserEnergy,
		"getCurrentExpedition": getCurrentExpedition,
		"createExpedition": createExpedition,
		"searchForVenariByName": searchForVenariByName,
		"searchForVenariByTier": searchForVenariByTier,
		"searchForVenariByAny": searchForVenariByAny,
		"getInventoryByName": getInventoryByName,
		"startBattle": startBattle,
		"battleAction": battleAction,
		"battleCatch": battleCatch,
		"buyItem": buyItem,
	}
	InvalidArgs = errors.New("invalid args")
	InvalidArgType = errors.New("invalid args type")
)

//////////////
//			//
//   USER   //
//			//
//////////////

// getUserEnergy[resVar]
func getUserEnergy(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 0 {
		return InvalidArgs
	}

	vm.SetValue(args[0], "")

	u, err := user.GetUser(t.Client)
	if err != nil {
		return err
	}

	vm.SetValue(args[0], u.Data.Energy.Amount)
	return nil
}

////////////////
//			  //
// EXPEDITION //
//			  //
////////////////

// getCurrentExpedition[area,resVar]
func getCurrentExpedition(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 1 {
		return InvalidArgs
	}

	vm.SetValue(args[1], "")

	area, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}

	areaStr, ok := area.(string)
	if !ok {
		return InvalidArgType
	}

	u, err := expeditions.GetExpeditions(areaStr, t.Client)
	if err != nil {
		return err
	}

	if len(*u) <= 0 {
		vm.SetValue(args[1], "")
	} else {
		vm.SetValue(args[1], (*u)[0].ID)
	}

	return nil
}

// createExpedition[area,baitId,resVar]
func createExpedition(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 2 {
		return InvalidArgs
	}

	vm.SetValue(args[2], "")

	area, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}

	areaStr, ok := area.(string)
	if !ok {
		return InvalidArgType
	}

	baitId, err := vm.GetValue(args[1])
	if err != nil {
		return err
	}

	baitIdStr, ok := baitId.(string)
	if !ok {
		return InvalidArgType
	}

	u, err := expeditions.CreateExpedition(areaStr, baitIdStr, t.Client)
	if err != nil {
		return err
	}

	if len(*u) <= 0 {
		vm.SetValue(args[2], "")
	} else {
		vm.SetValue(args[2], (*u)[0].ID)
	}

	return nil
}

////////////////
//			  //
//   SEARCH   //
//			  //
////////////////

// searchForVenariByName[area,name,resVar]
func searchForVenariByName(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 2 {
		return InvalidArgs
	}

	vm.SetValue(args[2], "")

	area, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}

	areaStr, ok := area.(string)
	if !ok {
		return InvalidArgType
	}

	name, err := vm.GetValue(args[1])
	if err != nil {
		return err
	}

	nameStr, ok := name.(string)
	if !ok {
		return InvalidArgType
	}

	currentExp, err := expeditions.GetExpeditions(areaStr, t.Client)
	if err != nil {
		return err
	}

	if len(*currentExp) <= 0 {
		return errors.New("no expedition found")
	}

	exp := (*currentExp)[0]

	for _, v := range exp.Spawns {
		if v.Venari.Name == nameStr {
			vm.SetValue(args[2], v.ID)
			break
		}
	}
	return nil
}

// searchForVenariByTier[area,tier,resVar]
func searchForVenariByTier(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 2 {
		return InvalidArgs
	}

	vm.SetValue(args[2], "")

	area, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}

	areaStr, ok := area.(string)
	if !ok {
		return InvalidArgType
	}

	tier, err := vm.GetValue(args[1])
	if err != nil {
		return err
	}

	tierStr, ok := tier.(string)
	if !ok {
		return InvalidArgType
	}

	currentExp, err := expeditions.GetExpeditions(areaStr, t.Client)
	if err != nil {
		return err
	}

	if len(*currentExp) <= 0 {
		return errors.New("no expedition found")
	}

	exp := (*currentExp)[0]

	for _, v := range exp.Spawns {
		if v.Venari.Tier == tierStr {
			vm.SetValue(args[2], v.ID)
			break
		}
	}
	return nil
}

// searchForVenariByAny[area,resVar]
func searchForVenariByAny(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 1 {
		return InvalidArgs
	}

	vm.SetValue(args[1], "")

	area, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}

	areaStr, ok := area.(string)
	if !ok {
		return InvalidArgType
	}

	currentExp, err := expeditions.GetExpeditions(areaStr, t.Client)
	if err != nil {
		return err
	}

	if len(*currentExp) <= 0 {
		return errors.New("no expedition found")
	}

	exp := (*currentExp)[0]

	for _, v := range exp.Spawns {
		vm.SetValue(args[1], v.ID)
		break
	}
	return nil
}

////////////////
//			  //
//  INVENTORY //
//			  //
////////////////

// getInventoryByName[name,resVar]
func getInventoryByName(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 1 {
		return InvalidArgs
	}

	vm.SetValue(args[1], "")

	name, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}

	nameStr, ok := name.(string)
	if !ok {
		return InvalidArgType
	}

	currentInv, err := inventory.GetInventory(t.Client)
	if err != nil {
		return err
	}

	for _, v := range *currentInv {
		if v.Item.Name == nameStr {
			vm.SetValue(args[1], v.ID)
			break
		}
	}

	return nil
}

////////////////
//			  //
//   BATTLE   //
//			  //
////////////////

// startBattle[expeditionId,venariId,baitId,rigId,resVar]
func startBattle(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 4 {
		return InvalidArgs
	}

	vm.SetValue(args[4], "")

	exp, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}
	expStr, ok := exp.(string)
	if !ok {
		return InvalidArgType
	}

	ven, err := vm.GetValue(args[1])
	if err != nil {
		return err
	}
	venStr, ok := ven.(string)
	if !ok {
		return InvalidArgType
	}

	bait, err := vm.GetValue(args[2])
	if err != nil {
		return err
	}
	baitStr, ok := bait.(string)
	if !ok {
		return InvalidArgType
	}

	rig, err := vm.GetValue(args[3])
	if err != nil {
		return err
	}
	rigStr, ok := rig.(string)
	if !ok {
		return InvalidArgType
	}

	start, err := battle.StartBattle(expStr, venStr, baitStr, rigStr, t.Client)
	if err != nil {
		return err
	}

	vm.SetValue(args[4], start.Venari.Name)

	return nil
}

// battleAction[expeditionId,baitId,rigId,action,resVar]
func battleAction(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 4 {
		return InvalidArgs
	}

	vm.SetValue(args[4], "")

	exp, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}
	expStr, ok := exp.(string)
	if !ok {
		return InvalidArgType
	}

	bait, err := vm.GetValue(args[1])
	if err != nil {
		return err
	}
	baitStr, ok := bait.(string)
	if !ok {
		return InvalidArgType
	}

	rig, err := vm.GetValue(args[2])
	if err != nil {
		return err
	}
	rigStr, ok := rig.(string)
	if !ok {
		return InvalidArgType
	}

	action, err := vm.GetValue(args[3])
	if err != nil {
		return err
	}
	actionStr, ok := action.(string)
	if !ok {
		return InvalidArgType
	}

	ac, ok := battle.BattleActionList[actionStr]
	if !ok {
		return errors.New("not a valid action type")
	}

	act, err := battle.BattleAction(expStr, baitStr, rigStr, ac, t.Client)
	if err != nil {
		return err
	}

	vm.SetValue(args[4], act.Success)

	return nil
}

// battleCatch[expeditionId,resVar]
func battleCatch(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 1 {
		return InvalidArgs
	}

	vm.SetValue(args[1], "")

	exp, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}
	expStr, ok := exp.(string)
	if !ok {
		return InvalidArgType
	}

	act, err := battle.BattleCatch(expStr, t.Client)
	if err != nil {
		return err
	}

	if act.Success {
		vm.SetValue(args[1], act.Rewards[0].Amount)
	} else {
		vm.SetValue(args[1], 0)
	}

	return nil
}

////////////////
//			  //
//    SHOP    //
//			  //
////////////////

// buyItem[itemId,amount,resVar]
func buyItem(vm *InterpreterTask, t *strats.VenariTask, args []string) error {
	if len(args) <= 2 {
		return InvalidArgs
	}

	vm.SetValue(args[2], "")

	item, err := vm.GetValue(args[0])
	if err != nil {
		return err
	}

	itemStr, ok := item.(string)
	if !ok {
		return InvalidArgType
	}

	amount, err := vm.GetValue(args[1])
	if err != nil {
		return err
	}

	amountStr, ok := amount.(string)
	if !ok {
		return InvalidArgType
	}

	amountNum, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	items, err := user.BuyItem(itemStr, amountNum, t.Client)
	if err != nil {
		return err
	}

	if strings.Contains(items.Message, "success") {
		vm.SetValue(args[2], items.Message)
	}

	return nil
}