class CreatePublishers < ActiveRecord::Migration[5.2]
  def up
    create_table :publishers do |t|
      t.string :name
      t.string :address
      t.string :phone_number
      t.timestamps
    end
  end

  def down 
    drop_table :publishers
  end
end
